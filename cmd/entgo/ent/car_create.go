// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/zrma/1d1c/cmd/entgo/ent/car"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	model         *string
	registered_at *time.Time
}

// SetModel sets the model field.
func (cc *CarCreate) SetModel(s string) *CarCreate {
	cc.model = &s
	return cc
}

// SetRegisteredAt sets the registered_at field.
func (cc *CarCreate) SetRegisteredAt(t time.Time) *CarCreate {
	cc.registered_at = &t
	return cc
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	if cc.model == nil {
		return nil, errors.New("ent: missing required field \"model\"")
	}
	if cc.registered_at == nil {
		return nil, errors.New("ent: missing required field \"registered_at\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	var (
		res sql.Result
		c   = &Car{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(car.Table).Default(cc.driver.Dialect())
	if value := cc.model; value != nil {
		builder.Set(car.FieldModel, *value)
		c.Model = *value
	}
	if value := cc.registered_at; value != nil {
		builder.Set(car.FieldRegisteredAt, *value)
		c.RegisteredAt = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
