// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// OccurrenceCreate is the builder for creating a Occurrence entity.
type OccurrenceCreate struct {
	config
	mutation *OccurrenceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetArtifactID sets the "artifact_id" field.
func (oc *OccurrenceCreate) SetArtifactID(i int) *OccurrenceCreate {
	oc.mutation.SetArtifactID(i)
	return oc
}

// SetJustification sets the "justification" field.
func (oc *OccurrenceCreate) SetJustification(s string) *OccurrenceCreate {
	oc.mutation.SetJustification(s)
	return oc
}

// SetOrigin sets the "origin" field.
func (oc *OccurrenceCreate) SetOrigin(s string) *OccurrenceCreate {
	oc.mutation.SetOrigin(s)
	return oc
}

// SetCollector sets the "collector" field.
func (oc *OccurrenceCreate) SetCollector(s string) *OccurrenceCreate {
	oc.mutation.SetCollector(s)
	return oc
}

// SetSourceID sets the "source_id" field.
func (oc *OccurrenceCreate) SetSourceID(i int) *OccurrenceCreate {
	oc.mutation.SetSourceID(i)
	return oc
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (oc *OccurrenceCreate) SetNillableSourceID(i *int) *OccurrenceCreate {
	if i != nil {
		oc.SetSourceID(*i)
	}
	return oc
}

// SetPackageID sets the "package_id" field.
func (oc *OccurrenceCreate) SetPackageID(i int) *OccurrenceCreate {
	oc.mutation.SetPackageID(i)
	return oc
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (oc *OccurrenceCreate) SetNillablePackageID(i *int) *OccurrenceCreate {
	if i != nil {
		oc.SetPackageID(*i)
	}
	return oc
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (oc *OccurrenceCreate) SetArtifact(a *Artifact) *OccurrenceCreate {
	return oc.SetArtifactID(a.ID)
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (oc *OccurrenceCreate) SetPackage(p *PackageVersion) *OccurrenceCreate {
	return oc.SetPackageID(p.ID)
}

// SetSource sets the "source" edge to the SourceName entity.
func (oc *OccurrenceCreate) SetSource(s *SourceName) *OccurrenceCreate {
	return oc.SetSourceID(s.ID)
}

// Mutation returns the OccurrenceMutation object of the builder.
func (oc *OccurrenceCreate) Mutation() *OccurrenceMutation {
	return oc.mutation
}

// Save creates the Occurrence in the database.
func (oc *OccurrenceCreate) Save(ctx context.Context) (*Occurrence, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OccurrenceCreate) SaveX(ctx context.Context) *Occurrence {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OccurrenceCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OccurrenceCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OccurrenceCreate) check() error {
	if _, ok := oc.mutation.ArtifactID(); !ok {
		return &ValidationError{Name: "artifact_id", err: errors.New(`ent: missing required field "Occurrence.artifact_id"`)}
	}
	if _, ok := oc.mutation.Justification(); !ok {
		return &ValidationError{Name: "justification", err: errors.New(`ent: missing required field "Occurrence.justification"`)}
	}
	if _, ok := oc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Occurrence.origin"`)}
	}
	if _, ok := oc.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "Occurrence.collector"`)}
	}
	if _, ok := oc.mutation.ArtifactID(); !ok {
		return &ValidationError{Name: "artifact", err: errors.New(`ent: missing required edge "Occurrence.artifact"`)}
	}
	return nil
}

func (oc *OccurrenceCreate) sqlSave(ctx context.Context) (*Occurrence, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OccurrenceCreate) createSpec() (*Occurrence, *sqlgraph.CreateSpec) {
	var (
		_node = &Occurrence{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(occurrence.Table, sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeInt))
	)
	_spec.OnConflict = oc.conflict
	if value, ok := oc.mutation.Justification(); ok {
		_spec.SetField(occurrence.FieldJustification, field.TypeString, value)
		_node.Justification = value
	}
	if value, ok := oc.mutation.Origin(); ok {
		_spec.SetField(occurrence.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := oc.mutation.Collector(); ok {
		_spec.SetField(occurrence.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if nodes := oc.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.ArtifactTable,
			Columns: []string{occurrence.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArtifactID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.PackageTable,
			Columns: []string{occurrence.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PackageID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.SourceTable,
			Columns: []string{occurrence.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Occurrence.Create().
//		SetArtifactID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OccurrenceUpsert) {
//			SetArtifactID(v+v).
//		}).
//		Exec(ctx)
func (oc *OccurrenceCreate) OnConflict(opts ...sql.ConflictOption) *OccurrenceUpsertOne {
	oc.conflict = opts
	return &OccurrenceUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oc *OccurrenceCreate) OnConflictColumns(columns ...string) *OccurrenceUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OccurrenceUpsertOne{
		create: oc,
	}
}

type (
	// OccurrenceUpsertOne is the builder for "upsert"-ing
	//  one Occurrence node.
	OccurrenceUpsertOne struct {
		create *OccurrenceCreate
	}

	// OccurrenceUpsert is the "OnConflict" setter.
	OccurrenceUpsert struct {
		*sql.UpdateSet
	}
)

// SetArtifactID sets the "artifact_id" field.
func (u *OccurrenceUpsert) SetArtifactID(v int) *OccurrenceUpsert {
	u.Set(occurrence.FieldArtifactID, v)
	return u
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdateArtifactID() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldArtifactID)
	return u
}

// SetJustification sets the "justification" field.
func (u *OccurrenceUpsert) SetJustification(v string) *OccurrenceUpsert {
	u.Set(occurrence.FieldJustification, v)
	return u
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdateJustification() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldJustification)
	return u
}

// SetOrigin sets the "origin" field.
func (u *OccurrenceUpsert) SetOrigin(v string) *OccurrenceUpsert {
	u.Set(occurrence.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdateOrigin() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *OccurrenceUpsert) SetCollector(v string) *OccurrenceUpsert {
	u.Set(occurrence.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdateCollector() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldCollector)
	return u
}

// SetSourceID sets the "source_id" field.
func (u *OccurrenceUpsert) SetSourceID(v int) *OccurrenceUpsert {
	u.Set(occurrence.FieldSourceID, v)
	return u
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdateSourceID() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldSourceID)
	return u
}

// ClearSourceID clears the value of the "source_id" field.
func (u *OccurrenceUpsert) ClearSourceID() *OccurrenceUpsert {
	u.SetNull(occurrence.FieldSourceID)
	return u
}

// SetPackageID sets the "package_id" field.
func (u *OccurrenceUpsert) SetPackageID(v int) *OccurrenceUpsert {
	u.Set(occurrence.FieldPackageID, v)
	return u
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *OccurrenceUpsert) UpdatePackageID() *OccurrenceUpsert {
	u.SetExcluded(occurrence.FieldPackageID)
	return u
}

// ClearPackageID clears the value of the "package_id" field.
func (u *OccurrenceUpsert) ClearPackageID() *OccurrenceUpsert {
	u.SetNull(occurrence.FieldPackageID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OccurrenceUpsertOne) UpdateNewValues() *OccurrenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OccurrenceUpsertOne) Ignore() *OccurrenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OccurrenceUpsertOne) DoNothing() *OccurrenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OccurrenceCreate.OnConflict
// documentation for more info.
func (u *OccurrenceUpsertOne) Update(set func(*OccurrenceUpsert)) *OccurrenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OccurrenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetArtifactID sets the "artifact_id" field.
func (u *OccurrenceUpsertOne) SetArtifactID(v int) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetArtifactID(v)
	})
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdateArtifactID() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateArtifactID()
	})
}

// SetJustification sets the "justification" field.
func (u *OccurrenceUpsertOne) SetJustification(v string) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdateJustification() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *OccurrenceUpsertOne) SetOrigin(v string) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdateOrigin() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *OccurrenceUpsertOne) SetCollector(v string) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdateCollector() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateCollector()
	})
}

// SetSourceID sets the "source_id" field.
func (u *OccurrenceUpsertOne) SetSourceID(v int) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdateSourceID() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateSourceID()
	})
}

// ClearSourceID clears the value of the "source_id" field.
func (u *OccurrenceUpsertOne) ClearSourceID() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.ClearSourceID()
	})
}

// SetPackageID sets the "package_id" field.
func (u *OccurrenceUpsertOne) SetPackageID(v int) *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *OccurrenceUpsertOne) UpdatePackageID() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdatePackageID()
	})
}

// ClearPackageID clears the value of the "package_id" field.
func (u *OccurrenceUpsertOne) ClearPackageID() *OccurrenceUpsertOne {
	return u.Update(func(s *OccurrenceUpsert) {
		s.ClearPackageID()
	})
}

// Exec executes the query.
func (u *OccurrenceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OccurrenceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OccurrenceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OccurrenceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OccurrenceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OccurrenceCreateBulk is the builder for creating many Occurrence entities in bulk.
type OccurrenceCreateBulk struct {
	config
	builders []*OccurrenceCreate
	conflict []sql.ConflictOption
}

// Save creates the Occurrence entities in the database.
func (ocb *OccurrenceCreateBulk) Save(ctx context.Context) ([]*Occurrence, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Occurrence, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OccurrenceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OccurrenceCreateBulk) SaveX(ctx context.Context) []*Occurrence {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OccurrenceCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OccurrenceCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Occurrence.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OccurrenceUpsert) {
//			SetArtifactID(v+v).
//		}).
//		Exec(ctx)
func (ocb *OccurrenceCreateBulk) OnConflict(opts ...sql.ConflictOption) *OccurrenceUpsertBulk {
	ocb.conflict = opts
	return &OccurrenceUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ocb *OccurrenceCreateBulk) OnConflictColumns(columns ...string) *OccurrenceUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OccurrenceUpsertBulk{
		create: ocb,
	}
}

// OccurrenceUpsertBulk is the builder for "upsert"-ing
// a bulk of Occurrence nodes.
type OccurrenceUpsertBulk struct {
	create *OccurrenceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OccurrenceUpsertBulk) UpdateNewValues() *OccurrenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Occurrence.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OccurrenceUpsertBulk) Ignore() *OccurrenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OccurrenceUpsertBulk) DoNothing() *OccurrenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OccurrenceCreateBulk.OnConflict
// documentation for more info.
func (u *OccurrenceUpsertBulk) Update(set func(*OccurrenceUpsert)) *OccurrenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OccurrenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetArtifactID sets the "artifact_id" field.
func (u *OccurrenceUpsertBulk) SetArtifactID(v int) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetArtifactID(v)
	})
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdateArtifactID() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateArtifactID()
	})
}

// SetJustification sets the "justification" field.
func (u *OccurrenceUpsertBulk) SetJustification(v string) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdateJustification() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *OccurrenceUpsertBulk) SetOrigin(v string) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdateOrigin() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *OccurrenceUpsertBulk) SetCollector(v string) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdateCollector() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateCollector()
	})
}

// SetSourceID sets the "source_id" field.
func (u *OccurrenceUpsertBulk) SetSourceID(v int) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdateSourceID() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdateSourceID()
	})
}

// ClearSourceID clears the value of the "source_id" field.
func (u *OccurrenceUpsertBulk) ClearSourceID() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.ClearSourceID()
	})
}

// SetPackageID sets the "package_id" field.
func (u *OccurrenceUpsertBulk) SetPackageID(v int) *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *OccurrenceUpsertBulk) UpdatePackageID() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.UpdatePackageID()
	})
}

// ClearPackageID clears the value of the "package_id" field.
func (u *OccurrenceUpsertBulk) ClearPackageID() *OccurrenceUpsertBulk {
	return u.Update(func(s *OccurrenceUpsert) {
		s.ClearPackageID()
	})
}

// Exec executes the query.
func (u *OccurrenceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OccurrenceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OccurrenceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OccurrenceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
