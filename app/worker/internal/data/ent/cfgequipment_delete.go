// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"harvest/app/worker/internal/data/ent/cfgequipment"
	"harvest/app/worker/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CfgEquipmentDelete is the builder for deleting a CfgEquipment entity.
type CfgEquipmentDelete struct {
	config
	hooks    []Hook
	mutation *CfgEquipmentMutation
}

// Where appends a list predicates to the CfgEquipmentDelete builder.
func (ced *CfgEquipmentDelete) Where(ps ...predicate.CfgEquipment) *CfgEquipmentDelete {
	ced.mutation.Where(ps...)
	return ced
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ced *CfgEquipmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ced.hooks) == 0 {
		affected, err = ced.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CfgEquipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ced.mutation = mutation
			affected, err = ced.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ced.hooks) - 1; i >= 0; i-- {
			if ced.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ced.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ced.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ced *CfgEquipmentDelete) ExecX(ctx context.Context) int {
	n, err := ced.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ced *CfgEquipmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: cfgequipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cfgequipment.FieldID,
			},
		},
	}
	if ps := ced.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ced.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CfgEquipmentDeleteOne is the builder for deleting a single CfgEquipment entity.
type CfgEquipmentDeleteOne struct {
	ced *CfgEquipmentDelete
}

// Exec executes the deletion query.
func (cedo *CfgEquipmentDeleteOne) Exec(ctx context.Context) error {
	n, err := cedo.ced.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cfgequipment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cedo *CfgEquipmentDeleteOne) ExecX(ctx context.Context) {
	cedo.ced.ExecX(ctx)
}