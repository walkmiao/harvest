// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"harvest/app/worker/internal/data/ent/migrate"

	"harvest/app/worker/internal/data/ent/cfgequipment"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CfgEquipment is the client for interacting with the CfgEquipment builders.
	CfgEquipment *CfgEquipmentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CfgEquipment = NewCfgEquipmentClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		CfgEquipment: NewCfgEquipmentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		CfgEquipment: NewCfgEquipmentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CfgEquipment.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CfgEquipment.Use(hooks...)
}

// CfgEquipmentClient is a client for the CfgEquipment schema.
type CfgEquipmentClient struct {
	config
}

// NewCfgEquipmentClient returns a client for the CfgEquipment from the given config.
func NewCfgEquipmentClient(c config) *CfgEquipmentClient {
	return &CfgEquipmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cfgequipment.Hooks(f(g(h())))`.
func (c *CfgEquipmentClient) Use(hooks ...Hook) {
	c.hooks.CfgEquipment = append(c.hooks.CfgEquipment, hooks...)
}

// Create returns a builder for creating a CfgEquipment entity.
func (c *CfgEquipmentClient) Create() *CfgEquipmentCreate {
	mutation := newCfgEquipmentMutation(c.config, OpCreate)
	return &CfgEquipmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CfgEquipment entities.
func (c *CfgEquipmentClient) CreateBulk(builders ...*CfgEquipmentCreate) *CfgEquipmentCreateBulk {
	return &CfgEquipmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CfgEquipment.
func (c *CfgEquipmentClient) Update() *CfgEquipmentUpdate {
	mutation := newCfgEquipmentMutation(c.config, OpUpdate)
	return &CfgEquipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CfgEquipmentClient) UpdateOne(ce *CfgEquipment) *CfgEquipmentUpdateOne {
	mutation := newCfgEquipmentMutation(c.config, OpUpdateOne, withCfgEquipment(ce))
	return &CfgEquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CfgEquipmentClient) UpdateOneID(id int) *CfgEquipmentUpdateOne {
	mutation := newCfgEquipmentMutation(c.config, OpUpdateOne, withCfgEquipmentID(id))
	return &CfgEquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CfgEquipment.
func (c *CfgEquipmentClient) Delete() *CfgEquipmentDelete {
	mutation := newCfgEquipmentMutation(c.config, OpDelete)
	return &CfgEquipmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CfgEquipmentClient) DeleteOne(ce *CfgEquipment) *CfgEquipmentDeleteOne {
	return c.DeleteOneID(ce.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CfgEquipmentClient) DeleteOneID(id int) *CfgEquipmentDeleteOne {
	builder := c.Delete().Where(cfgequipment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CfgEquipmentDeleteOne{builder}
}

// Query returns a query builder for CfgEquipment.
func (c *CfgEquipmentClient) Query() *CfgEquipmentQuery {
	return &CfgEquipmentQuery{
		config: c.config,
	}
}

// Get returns a CfgEquipment entity by its id.
func (c *CfgEquipmentClient) Get(ctx context.Context, id int) (*CfgEquipment, error) {
	return c.Query().Where(cfgequipment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CfgEquipmentClient) GetX(ctx context.Context, id int) *CfgEquipment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CfgEquipmentClient) Hooks() []Hook {
	return c.hooks.CfgEquipment
}
