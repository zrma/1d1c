// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/zrma/1d1c/cmd/entgo/ent/group"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	name *string
}

// SetName sets the name field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.name = &s
	return gc
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if gc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := group.NameValidator(*gc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	return gc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	var (
		res sql.Result
		gr  = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(group.Table).Default(gc.driver.Dialect())
	if value := gc.name; value != nil {
		builder.Set(group.FieldName, *value)
		gr.Name = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
