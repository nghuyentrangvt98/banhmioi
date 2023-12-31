// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nghuyentrangvt98/banhmioi/ent/cart"
	"github.com/nghuyentrangvt98/banhmioi/ent/predicate"
	"github.com/nghuyentrangvt98/banhmioi/ent/product"
	"github.com/nghuyentrangvt98/banhmioi/ent/user"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetVariation sets the "variation" field.
func (cu *CartUpdate) SetVariation(m map[string]interface{}) *CartUpdate {
	cu.mutation.SetVariation(m)
	return cu
}

// SetQty sets the "qty" field.
func (cu *CartUpdate) SetQty(i int32) *CartUpdate {
	cu.mutation.ResetQty()
	cu.mutation.SetQty(i)
	return cu
}

// AddQty adds i to the "qty" field.
func (cu *CartUpdate) AddQty(i int32) *CartUpdate {
	cu.mutation.AddQty(i)
	return cu
}

// SetPrice sets the "price" field.
func (cu *CartUpdate) SetPrice(i int64) *CartUpdate {
	cu.mutation.ResetPrice()
	cu.mutation.SetPrice(i)
	return cu
}

// AddPrice adds i to the "price" field.
func (cu *CartUpdate) AddPrice(i int64) *CartUpdate {
	cu.mutation.AddPrice(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CartUpdate) SetStatus(c cart.Status) *CartUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNote sets the "note" field.
func (cu *CartUpdate) SetNote(s string) *CartUpdate {
	cu.mutation.SetNote(s)
	return cu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (cu *CartUpdate) SetNillableNote(s *string) *CartUpdate {
	if s != nil {
		cu.SetNote(*s)
	}
	return cu
}

// ClearNote clears the value of the "note" field.
func (cu *CartUpdate) ClearNote() *CartUpdate {
	cu.mutation.ClearNote()
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CartUpdate) SetUserID(id int) *CartUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cu *CartUpdate) SetNillableUserID(id *int) *CartUpdate {
	if id != nil {
		cu = cu.SetUserID(*id)
	}
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CartUpdate) SetUser(u *User) *CartUpdate {
	return cu.SetUserID(u.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (cu *CartUpdate) SetProductID(id int) *CartUpdate {
	cu.mutation.SetProductID(id)
	return cu
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (cu *CartUpdate) SetNillableProductID(id *int) *CartUpdate {
	if id != nil {
		cu = cu.SetProductID(*id)
	}
	return cu
}

// SetProduct sets the "product" edge to the Product entity.
func (cu *CartUpdate) SetProduct(p *Product) *CartUpdate {
	return cu.SetProductID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CartUpdate) ClearUser() *CartUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearProduct clears the "product" edge to the Product entity.
func (cu *CartUpdate) ClearProduct() *CartUpdate {
	cu.mutation.ClearProduct()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CartUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := cart.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Cart.status": %w`, err)}
		}
	}
	return nil
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Variation(); ok {
		_spec.SetField(cart.FieldVariation, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.Qty(); ok {
		_spec.SetField(cart.FieldQty, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedQty(); ok {
		_spec.AddField(cart.FieldQty, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.Price(); ok {
		_spec.SetField(cart.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedPrice(); ok {
		_spec.AddField(cart.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(cart.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Note(); ok {
		_spec.SetField(cart.FieldNote, field.TypeString, value)
	}
	if cu.mutation.NoteCleared() {
		_spec.ClearField(cart.FieldNote, field.TypeString)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.ProductTable,
			Columns: []string{cart.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.ProductTable,
			Columns: []string{cart.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetVariation sets the "variation" field.
func (cuo *CartUpdateOne) SetVariation(m map[string]interface{}) *CartUpdateOne {
	cuo.mutation.SetVariation(m)
	return cuo
}

// SetQty sets the "qty" field.
func (cuo *CartUpdateOne) SetQty(i int32) *CartUpdateOne {
	cuo.mutation.ResetQty()
	cuo.mutation.SetQty(i)
	return cuo
}

// AddQty adds i to the "qty" field.
func (cuo *CartUpdateOne) AddQty(i int32) *CartUpdateOne {
	cuo.mutation.AddQty(i)
	return cuo
}

// SetPrice sets the "price" field.
func (cuo *CartUpdateOne) SetPrice(i int64) *CartUpdateOne {
	cuo.mutation.ResetPrice()
	cuo.mutation.SetPrice(i)
	return cuo
}

// AddPrice adds i to the "price" field.
func (cuo *CartUpdateOne) AddPrice(i int64) *CartUpdateOne {
	cuo.mutation.AddPrice(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CartUpdateOne) SetStatus(c cart.Status) *CartUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNote sets the "note" field.
func (cuo *CartUpdateOne) SetNote(s string) *CartUpdateOne {
	cuo.mutation.SetNote(s)
	return cuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableNote(s *string) *CartUpdateOne {
	if s != nil {
		cuo.SetNote(*s)
	}
	return cuo
}

// ClearNote clears the value of the "note" field.
func (cuo *CartUpdateOne) ClearNote() *CartUpdateOne {
	cuo.mutation.ClearNote()
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CartUpdateOne) SetUserID(id int) *CartUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUserID(id *int) *CartUpdateOne {
	if id != nil {
		cuo = cuo.SetUserID(*id)
	}
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CartUpdateOne) SetUser(u *User) *CartUpdateOne {
	return cuo.SetUserID(u.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (cuo *CartUpdateOne) SetProductID(id int) *CartUpdateOne {
	cuo.mutation.SetProductID(id)
	return cuo
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableProductID(id *int) *CartUpdateOne {
	if id != nil {
		cuo = cuo.SetProductID(*id)
	}
	return cuo
}

// SetProduct sets the "product" edge to the Product entity.
func (cuo *CartUpdateOne) SetProduct(p *Product) *CartUpdateOne {
	return cuo.SetProductID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CartUpdateOne) ClearUser() *CartUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearProduct clears the "product" edge to the Product entity.
func (cuo *CartUpdateOne) ClearProduct() *CartUpdateOne {
	cuo.mutation.ClearProduct()
	return cuo
}

// Where appends a list predicates to the CartUpdate builder.
func (cuo *CartUpdateOne) Where(ps ...predicate.Cart) *CartUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CartUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := cart.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Cart.status": %w`, err)}
		}
	}
	return nil
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Variation(); ok {
		_spec.SetField(cart.FieldVariation, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.Qty(); ok {
		_spec.SetField(cart.FieldQty, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedQty(); ok {
		_spec.AddField(cart.FieldQty, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.Price(); ok {
		_spec.SetField(cart.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedPrice(); ok {
		_spec.AddField(cart.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(cart.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Note(); ok {
		_spec.SetField(cart.FieldNote, field.TypeString, value)
	}
	if cuo.mutation.NoteCleared() {
		_spec.ClearField(cart.FieldNote, field.TypeString)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.ProductTable,
			Columns: []string{cart.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cart.ProductTable,
			Columns: []string{cart.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
