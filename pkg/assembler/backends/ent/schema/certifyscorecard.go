//
// Copyright 2023 The GUAC Authors.
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

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CertifyScorecard holds the schema definition for the CertifyScorecard entity.
type CertifyScorecard struct {
	ent.Schema
}

// Fields of the CertifyScorecard.
func (CertifyScorecard) Fields() []ent.Field {
	return []ent.Field{
		field.Int("source_id"),
		field.Int("scorecard_id"),
	}
}

// Edges of the CertifyScorecard.
func (CertifyScorecard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scorecard", Scorecard.Type).Unique().Required().Ref("certifications").Field("scorecard_id"),
		edge.To("source", SourceName.Type).Unique().Required().Field("source_id"),
	}
}
func (CertifyScorecard) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_id", "scorecard_id").Unique(),
	}
}
