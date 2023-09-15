//
// Copyright 2022 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cyclonedx

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"

	"github.com/guacsec/guac/pkg/assembler"
	model "github.com/guacsec/guac/pkg/assembler/clients/generated"
	asmhelpers "github.com/guacsec/guac/pkg/assembler/helpers"
	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/guacsec/guac/pkg/ingestor/parser/common"
	"github.com/guacsec/guac/pkg/logging"
)

const topCdxPurlGuac string = "pkg:guac/cdx/"

var vexStatusMap = map[cdx.ImpactAnalysisState]model.VexStatus{
	"resolved":     model.VexStatusFixed,
	"exploitable":  model.VexStatusAffected,
	"in_triage":    model.VexStatusUnderInvestigation,
	"not_affected": model.VexStatusNotAffected,
}

var justificationsMap = map[cdx.ImpactAnalysisJustification]model.VexJustification{
	"code_not_present":   model.VexJustificationVulnerableCodeNotPresent,
	"code_not_reachable": model.VexJustificationVulnerableCodeNotInExecutePath,
}

type cyclonedxParser struct {
	doc               *processor.Document
	packagePackages   map[string][]*model.PkgInputSpec
	packageArtifacts  map[string][]*model.ArtifactInputSpec
	identifierStrings *common.IdentifierStrings
	cdxBom            *cdx.BOM
	vulnData          vulnData
}

type vulnData struct {
	vulnMetadata []assembler.VulnMetadataIngest
	certifyVuln  []assembler.CertifyVulnIngest
	vex          []assembler.VexIngest
}

func NewCycloneDXParser() common.DocumentParser {
	return &cyclonedxParser{
		packagePackages:   map[string][]*model.PkgInputSpec{},
		packageArtifacts:  map[string][]*model.ArtifactInputSpec{},
		identifierStrings: &common.IdentifierStrings{},
	}
}

// Parse breaks out the document into the graph components
func (c *cyclonedxParser) Parse(ctx context.Context, doc *processor.Document) error {
	c.doc = doc
	cdxBom, err := ParseCycloneDXBOM(doc)
	if err != nil {
		return fmt.Errorf("failed to parse cyclonedx BOM: %w", err)
	}
	c.cdxBom = cdxBom
	if err := c.getTopLevelPackage(cdxBom); err != nil {
		return err
	}
	if err := c.getPackages(cdxBom); err != nil {
		return err
	}
	if err := c.getVulnerabilities(ctx, cdxBom); err != nil {
		return err
	}

	return nil
}

// GetIdentities gets the identity node from the document if they exist
func (c *cyclonedxParser) GetIdentities(ctx context.Context) []common.TrustInformation {
	return nil
}

func (c *cyclonedxParser) getTopLevelPackage(cdxBom *cdx.BOM) error {
	if cdxBom.Metadata == nil {
		return nil
	}

	if cdxBom.Metadata.Component != nil {
		purl := cdxBom.Metadata.Component.PackageURL
		if cdxBom.Metadata.Component.PackageURL == "" {
			if cdxBom.Metadata.Component.Type == cdx.ComponentTypeContainer {
				purl = parseContainerType(cdxBom.Metadata.Component.Name, cdxBom.Metadata.Component.Version, true)
			} else if cdxBom.Metadata.Component.Type == cdx.ComponentTypeFile {
				// example: file type ("/home/work/test/build/webserver")
				purl = guacCDXFilePurl(cdxBom.Metadata.Component.Name, cdxBom.Metadata.Component.Version, true)
			} else {
				purl = guacCDXPkgPurl(cdxBom.Metadata.Component.Name, cdxBom.Metadata.Component.Version, "", true)
			}
		}

		topPackage, err := asmhelpers.PurlToPkg(purl)
		if err != nil {
			return err
		}
		c.identifierStrings.PurlStrings = append(c.identifierStrings.PurlStrings, purl)

		c.packagePackages[string(cdxBom.Metadata.Component.BOMRef)] = append(c.packagePackages[string(cdxBom.Metadata.Component.BOMRef)], topPackage)

		// if checksums exists create an artifact for each of them
		if cdxBom.Metadata.Component.Hashes != nil {
			for _, checksum := range *cdxBom.Metadata.Component.Hashes {
				artifact := &model.ArtifactInputSpec{
					Algorithm: strings.ToLower(string(checksum.Algorithm)),
					Digest:    checksum.Value,
				}
				c.packageArtifacts[string(cdxBom.Metadata.Component.BOMRef)] = append(c.packageArtifacts[string(cdxBom.Metadata.Component.BOMRef)], artifact)
			}
		}
		return nil
	} else {
		// currently GUAC does not support CycloneDX component field in metadata or the BOM ref being nil.
		// see https://github.com/guacsec/guac/issues/976 for more details.
		return fmt.Errorf("guac currently does not support CycloneDX component field in metadata or the BOM ref being nil. See issue #976 for more details")
	}
}

func parseContainerType(name string, version string, topLevel bool) string {
	splitImage := strings.Split(name, "/")
	splitTag := strings.Split(splitImage[len(splitImage)-1], ":")
	var repositoryURL string
	var tag string

	switch len(splitImage) {
	case 3:
		repositoryURL = splitImage[0] + "/" + splitImage[1] + "/" + splitTag[0]
	case 2:
		repositoryURL = splitImage[0] + "/" + splitTag[0]
	case 1:
		repositoryURL = splitImage[0]
	default:
		repositoryURL = ""
	}

	if len(splitTag) == 2 {
		tag = splitTag[1]
	}
	if repositoryURL != "" {
		return guacCDXPkgPurl(repositoryURL, version, tag, topLevel)
	} else {
		return guacCDXPkgPurl(name, version, tag, topLevel)
	}
}

func (c *cyclonedxParser) getPackages(cdxBom *cdx.BOM) error {
	if cdxBom.Components != nil {
		for _, comp := range *cdxBom.Components {
			// skipping over the "operating-system" type as it does not contain
			// the required purl for package node. Currently there is no use-case
			// to capture OS for GUAC.
			if comp.Type != cdx.ComponentTypeOS {
				purl := comp.PackageURL
				if purl == "" {
					if comp.Type == cdx.ComponentTypeContainer {
						purl = parseContainerType(comp.Name, comp.Version, false)
					} else if comp.Type == cdx.ComponentTypeFile {
						purl = guacCDXFilePurl(comp.Name, comp.Version, false)
					} else {
						purl = asmhelpers.GuacPkgPurl(comp.Name, &comp.Version)
					}
				}
				pkg, err := asmhelpers.PurlToPkg(purl)
				if err != nil {
					return err
				}
				c.packagePackages[string(comp.BOMRef)] = append(c.packagePackages[string(comp.BOMRef)], pkg)
				c.identifierStrings.PurlStrings = append(c.identifierStrings.PurlStrings, comp.PackageURL)

				// if checksums exists create an artifact for each of them
				if comp.Hashes != nil {
					for _, checksum := range *comp.Hashes {
						artifact := &model.ArtifactInputSpec{
							Algorithm: strings.ToLower(string(checksum.Algorithm)),
							Digest:    checksum.Value,
						}
						c.packageArtifacts[string(comp.BOMRef)] = append(c.packageArtifacts[string(comp.BOMRef)], artifact)
					}
				}
			}
		}
	}
	return nil
}

func ParseCycloneDXBOM(doc *processor.Document) (*cdx.BOM, error) {
	bom := cdx.BOM{}
	switch doc.Format {
	case processor.FormatJSON:
		if err := json.Unmarshal(doc.Blob, &bom); err != nil {
			return nil, err
		}
	case processor.FormatXML:
		if err := xml.Unmarshal(doc.Blob, &bom); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unrecognized CycloneDX format %s", doc.Format)
	}
	return &bom, nil
}

func (c *cyclonedxParser) GetIdentifiers(ctx context.Context) (*common.IdentifierStrings, error) {
	return c.identifierStrings, nil
}

func (c *cyclonedxParser) GetPredicates(ctx context.Context) *assembler.IngestPredicates {
	logger := logging.FromContext(ctx)
	preds := &assembler.IngestPredicates{}
	var toplevel []*model.PkgInputSpec

	if c.cdxBom.Metadata != nil && c.cdxBom.Metadata.Component != nil {
		toplevel = c.getPackageElement(string(c.cdxBom.Metadata.Component.BOMRef))
	}

	// adding top level package edge manually for all depends on package
	// TODO: This is not based on the relationship so that can be inaccurate (can capture both direct and in-direct)...Remove this and be done below by the *c.cdxBom.Dependencies?
	// see https://github.com/CycloneDX/specification/issues/33
	if toplevel != nil {
		preds.IsDependency = append(preds.IsDependency, common.CreateTopLevelIsDeps(toplevel[0], c.packagePackages, nil, "top-level package GUAC heuristic connecting to each file/package")...)
		preds.HasSBOM = append(preds.HasSBOM, common.CreateTopLevelHasSBOM(toplevel[0], c.doc))
	}

	for id := range c.packagePackages {
		for _, pkg := range c.packagePackages[id] {
			for _, art := range c.packageArtifacts[id] {
				preds.IsOccurrence = append(preds.IsOccurrence, assembler.IsOccurrenceIngest{
					Pkg:      pkg,
					Artifact: art,
					IsOccurrence: &model.IsOccurrenceInputSpec{
						Justification: "cdx package with checksum",
					},
				})
			}
		}
	}

	preds.Vex = c.vulnData.vex
	preds.VulnMetadata = c.vulnData.vulnMetadata
	preds.CertifyVuln = c.vulnData.certifyVuln
	if c.cdxBom.Dependencies == nil {
		return preds
	}

	for _, deps := range *c.cdxBom.Dependencies {
		currPkg, found := c.packagePackages[deps.Ref]
		if !found {
			continue
		}
		if reflect.DeepEqual(currPkg, toplevel) {
			continue
		}
		if deps.Dependencies != nil {
			for _, depPkg := range *deps.Dependencies {
				if depPkg, exist := c.packagePackages[depPkg]; exist {
					for _, packNode := range currPkg {
						p, err := common.GetIsDep(packNode, depPkg, []*model.PkgInputSpec{}, "CDX BOM Dependency")
						if err != nil {
							logger.Errorf("error generating CycloneDX edge %v", err)
							continue
						}
						if p != nil {
							preds.IsDependency = append(preds.IsDependency, *p)
						}
					}
				}
			}
		}
	}

	return preds
}

func (c *cyclonedxParser) getVulnerabilities(ctx context.Context, cdxBom *cdx.BOM) error {
	logger := logging.FromContext(ctx)
	if cdxBom.Vulnerabilities == nil {
		logger.Info("no vulnerabilities found in CycloneDX BOM")
		return nil
	}

	var status model.VexStatus
	var justification model.VexJustification
	var publishedTime time.Time
	for _, vulnerability := range *c.cdxBom.Vulnerabilities {
		vuln, err := asmhelpers.CreateVulnInput(vulnerability.ID)
		if err != nil {
			return err
		}

		if vexStatus, ok := vexStatusMap[vulnerability.Analysis.State]; ok {
			status = vexStatus
		} else {
			logger.Debugf("unknown vulnerability status %s", vulnerability.Analysis.State)
			status = "UNKNOWN"
		}

		if vexJustification, ok := justificationsMap[vulnerability.Analysis.Justification]; ok {
			justification = vexJustification
		} else {
			logger.Debugf("unknown vulnerability justification %s", vulnerability.Analysis.Justification)
			justification = "UNKNOWN"
		}

		time, err := time.Parse(time.RFC3339, vulnerability.Published)
		if err == nil {
			publishedTime = time
		}

		vd := model.VexStatementInputSpec{
			Status:           status,
			VexJustification: justification,
			KnownSince:       publishedTime,
			Statement:        vulnerability.Analysis.Detail,
			StatusNotes:      fmt.Sprintf("%s:%s", string(status), string(justification)),
		}

		for _, affect := range *vulnerability.Affects {
			vi := c.getAffectedPackages(ctx, vuln, vd, affect)
			if vi == nil {
				logger.Debugf("failed to get affected packages for vulnerability %s", vulnerability.ID)
				continue
			}
			c.vulnData.vex = append(c.vulnData.vex, *vi...)

			for _, v := range *vi {
				if status == model.VexStatusAffected || status == model.VexStatusUnderInvestigation {
					cv := assembler.CertifyVulnIngest{
						Vulnerability: vuln,
						VulnData: &model.ScanMetadataInput{
							TimeScanned: publishedTime,
						},
						Pkg: v.Pkg,
					}
					c.vulnData.certifyVuln = append(c.vulnData.certifyVuln, cv)
				}
			}
		}

		for _, vulnRating := range *vulnerability.Ratings {
			vm := assembler.VulnMetadataIngest{
				Vulnerability: vuln,
				VulnMetadata: &model.VulnerabilityMetadataInputSpec{
					ScoreType:  model.VulnerabilityScoreType(vulnRating.Method),
					ScoreValue: *vulnRating.Score,
					Timestamp:  publishedTime,
				},
			}
			c.vulnData.vulnMetadata = append(c.vulnData.vulnMetadata, vm)
		}
	}

	return nil
}

// Get package name and range versions to create package input spec for the affected packages.
func (c *cyclonedxParser) getAffectedPackages(ctx context.Context, vulnInput *model.VulnerabilityInputSpec, vexData model.VexStatementInputSpec, affectsObj cdx.Affects) *[]assembler.VexIngest {
	logger := logging.FromContext(ctx)
	var pkgRef string
	// TODO: retrieve purl from metadata if present - https://github.com/guacsec/guac/blob/main/pkg/ingestor/parser/cyclonedx/parser_cyclonedx.go#L76
	if affectsObj.Ref != "" {
		pkgRef = affectsObj.Ref
	} else {
		logger.Debugf("[cdx vex] package reference not found")
		return nil
	}

	// split ref using # as delimiter.
	pkgRefInfo := strings.Split(pkgRef, "#")
	if len(pkgRefInfo) != 2 {
		logger.Debugf("[cdx vex] malformed package reference: %q", affectsObj.Ref)
		return nil
	}
	pkgURL := pkgRefInfo[1]

	// multiple package versions do not exist, resolve to using ref directly.
	if affectsObj.Range == nil {
		pkg, err := asmhelpers.PurlToPkg(pkgURL)
		if err != nil {
			logger.Debugf("[cdx vex] unable to create package input spec: %v", err)
			return nil
		}

		c.identifierStrings.PurlStrings = append(c.identifierStrings.PurlStrings, pkgURL)
		return &[]assembler.VexIngest{{VexData: &vexData, Vulnerability: vulnInput, Pkg: pkg}}
	}

	// split pkgURL using @ as delimiter.
	pkgURLInfo := strings.Split(pkgURL, "@")
	if len(pkgURLInfo) != 2 {
		logger.Debugf("[cdx vex] malformed package url info: %q", pkgURL)
		return nil
	}

	pkgName := pkgURLInfo[0]
	var viList []assembler.VexIngest
	for _, affect := range *affectsObj.Range {
		// TODO: Handle package range versions (see - https://github.com/CycloneDX/bom-examples/blob/master/VEX/CISA-Use-Cases/Case-8/vex.json#L42)
		if affect.Version == "" {
			continue
		}
		vi := &assembler.VexIngest{
			VexData:       &vexData,
			Vulnerability: vulnInput,
		}

		pkg, err := asmhelpers.PurlToPkg(fmt.Sprintf("%s@%s", pkgName, affect.Version))
		if err != nil {
			logger.Debugf("[cdx vex] unable to create package input spec from purl: %v", err)
			return nil
		}
		vi.Pkg = pkg
		viList = append(viList, *vi)
		c.identifierStrings.PurlStrings = append(c.identifierStrings.PurlStrings, pkgURL)
	}

	return &viList
}

func (c *cyclonedxParser) getPackageElement(elementID string) []*model.PkgInputSpec {
	if packNode, ok := c.packagePackages[string(elementID)]; ok {
		return packNode
	}
	return nil
}

func guacCDXFilePurl(fileName string, version string, topLevel bool) string {
	escapedName := asmhelpers.SanitizeString(fileName)
	if topLevel {
		if version != "" {
			splitVersion := strings.Split(version, ":")
			if len(splitVersion) == 2 {
				s := fmt.Sprintf(topCdxPurlGuac+"%s:%s", strings.ToLower(splitVersion[0]), splitVersion[1])
				s += fmt.Sprintf("#%s", escapedName)
				return s
			}
		}
		return topCdxPurlGuac + escapedName
	} else {
		if version != "" {
			splitVersion := strings.Split(version, ":")
			if len(splitVersion) == 2 {
				return asmhelpers.GuacFilePurl(splitVersion[0], splitVersion[1], &escapedName)
			}
		}
		return asmhelpers.PurlFilesGuac + escapedName
	}
}

func guacCDXPkgPurl(componentName string, version string, tag string, topLevel bool) string {
	purl := ""
	typeNamespaceString := ""
	escapedName := asmhelpers.SanitizeString(componentName)
	if topLevel {
		typeNamespaceString = topCdxPurlGuac
	} else {
		typeNamespaceString = asmhelpers.PurlPkgGuac
	}
	if version != "" && tag != "" {
		purl = typeNamespaceString + escapedName + "@" + version + "?tag=" + tag
	} else if version != "" {
		purl = typeNamespaceString + escapedName + "@" + version
	} else if tag != "" {
		purl = typeNamespaceString + escapedName + "?tag=" + tag
	} else {
		purl = typeNamespaceString + escapedName
	}
	return purl
}
