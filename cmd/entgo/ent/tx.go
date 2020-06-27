// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect"
	"github.com/zrma/1d1go/cmd/entgo/ent/migrate"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	return tx.config.driver.(*txDriver).tx.Commit()
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	return tx.config.driver.(*txDriver).tx.Rollback()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	return &Client{
		config: tx.config,
		Schema: migrate.NewSchema(tx.driver),
		Car:    NewCarClient(tx.config),
		Group:  NewGroupClient(tx.config),
		User:   NewUserClient(tx.config),
	}
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Car.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
