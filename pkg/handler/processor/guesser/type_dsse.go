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

package guesser

import (
	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
)

type dsseTypeGuesser struct{}

func (_ *dsseTypeGuesser) GuessDocumentType(blob []byte, format processor.FormatType) processor.DocumentType {
	var envelope dsse.Envelope
	if json.Unmarshal(blob, &envelope) == nil && format == processor.FormatJSON {
		if envelope.Payload != "" && envelope.PayloadType != "" && len(envelope.Signatures) > 0 {
			return processor.DocumentDSSE
		}
	}
	return processor.DocumentUnknown
}
