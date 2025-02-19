// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifyvex"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// CertifyVexDelete is the builder for deleting a CertifyVex entity.
type CertifyVexDelete struct {
	config
	hooks    []Hook
	mutation *CertifyVexMutation
}

// Where appends a list predicates to the CertifyVexDelete builder.
func (cvd *CertifyVexDelete) Where(ps ...predicate.CertifyVex) *CertifyVexDelete {
	cvd.mutation.Where(ps...)
	return cvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cvd *CertifyVexDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cvd.sqlExec, cvd.mutation, cvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cvd *CertifyVexDelete) ExecX(ctx context.Context) int {
	n, err := cvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cvd *CertifyVexDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(certifyvex.Table, sqlgraph.NewFieldSpec(certifyvex.FieldID, field.TypeInt))
	if ps := cvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cvd.mutation.done = true
	return affected, err
}

// CertifyVexDeleteOne is the builder for deleting a single CertifyVex entity.
type CertifyVexDeleteOne struct {
	cvd *CertifyVexDelete
}

// Where appends a list predicates to the CertifyVexDelete builder.
func (cvdo *CertifyVexDeleteOne) Where(ps ...predicate.CertifyVex) *CertifyVexDeleteOne {
	cvdo.cvd.mutation.Where(ps...)
	return cvdo
}

// Exec executes the deletion query.
func (cvdo *CertifyVexDeleteOne) Exec(ctx context.Context) error {
	n, err := cvdo.cvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{certifyvex.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cvdo *CertifyVexDeleteOne) ExecX(ctx context.Context) {
	if err := cvdo.Exec(ctx); err != nil {
		panic(err)
	}
}
