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
	"github.com/nghuyentrangvt98/banhmioi/ent/order"
	"github.com/nghuyentrangvt98/banhmioi/ent/predicate"
	"github.com/nghuyentrangvt98/banhmioi/ent/user"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OrderUpdate) SetName(s string) *OrderUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetPhone sets the "phone" field.
func (ou *OrderUpdate) SetPhone(s string) *OrderUpdate {
	ou.mutation.SetPhone(s)
	return ou
}

// SetAddress sets the "address" field.
func (ou *OrderUpdate) SetAddress(s string) *OrderUpdate {
	ou.mutation.SetAddress(s)
	return ou
}

// SetNote sets the "note" field.
func (ou *OrderUpdate) SetNote(s string) *OrderUpdate {
	ou.mutation.SetNote(s)
	return ou
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableNote(s *string) *OrderUpdate {
	if s != nil {
		ou.SetNote(*s)
	}
	return ou
}

// ClearNote clears the value of the "note" field.
func (ou *OrderUpdate) ClearNote() *OrderUpdate {
	ou.mutation.ClearNote()
	return ou
}

// SetTotal sets the "total" field.
func (ou *OrderUpdate) SetTotal(f float32) *OrderUpdate {
	ou.mutation.ResetTotal()
	ou.mutation.SetTotal(f)
	return ou
}

// AddTotal adds f to the "total" field.
func (ou *OrderUpdate) AddTotal(f float32) *OrderUpdate {
	ou.mutation.AddTotal(f)
	return ou
}

// SetDiscount sets the "discount" field.
func (ou *OrderUpdate) SetDiscount(f float32) *OrderUpdate {
	ou.mutation.ResetDiscount()
	ou.mutation.SetDiscount(f)
	return ou
}

// AddDiscount adds f to the "discount" field.
func (ou *OrderUpdate) AddDiscount(f float32) *OrderUpdate {
	ou.mutation.AddDiscount(f)
	return ou
}

// AddCartIDs adds the "cart" edge to the Cart entity by IDs.
func (ou *OrderUpdate) AddCartIDs(ids ...int) *OrderUpdate {
	ou.mutation.AddCartIDs(ids...)
	return ou
}

// AddCart adds the "cart" edges to the Cart entity.
func (ou *OrderUpdate) AddCart(c ...*Cart) *OrderUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ou.AddCartIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OrderUpdate) SetUserID(id int) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetUserID(*id)
	}
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearCart clears all "cart" edges to the Cart entity.
func (ou *OrderUpdate) ClearCart() *OrderUpdate {
	ou.mutation.ClearCart()
	return ou
}

// RemoveCartIDs removes the "cart" edge to Cart entities by IDs.
func (ou *OrderUpdate) RemoveCartIDs(ids ...int) *OrderUpdate {
	ou.mutation.RemoveCartIDs(ids...)
	return ou
}

// RemoveCart removes "cart" edges to Cart entities.
func (ou *OrderUpdate) RemoveCart(c ...*Cart) *OrderUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ou.RemoveCartIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(order.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Phone(); ok {
		_spec.SetField(order.FieldPhone, field.TypeString, value)
	}
	if value, ok := ou.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
	}
	if value, ok := ou.mutation.Note(); ok {
		_spec.SetField(order.FieldNote, field.TypeString, value)
	}
	if ou.mutation.NoteCleared() {
		_spec.ClearField(order.FieldNote, field.TypeString)
	}
	if value, ok := ou.mutation.Total(); ok {
		_spec.SetField(order.FieldTotal, field.TypeFloat32, value)
	}
	if value, ok := ou.mutation.AddedTotal(); ok {
		_spec.AddField(order.FieldTotal, field.TypeFloat32, value)
	}
	if value, ok := ou.mutation.Discount(); ok {
		_spec.SetField(order.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ou.mutation.AddedDiscount(); ok {
		_spec.AddField(order.FieldDiscount, field.TypeFloat32, value)
	}
	if ou.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedCartIDs(); len(nodes) > 0 && !ou.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetName sets the "name" field.
func (ouo *OrderUpdateOne) SetName(s string) *OrderUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetPhone sets the "phone" field.
func (ouo *OrderUpdateOne) SetPhone(s string) *OrderUpdateOne {
	ouo.mutation.SetPhone(s)
	return ouo
}

// SetAddress sets the "address" field.
func (ouo *OrderUpdateOne) SetAddress(s string) *OrderUpdateOne {
	ouo.mutation.SetAddress(s)
	return ouo
}

// SetNote sets the "note" field.
func (ouo *OrderUpdateOne) SetNote(s string) *OrderUpdateOne {
	ouo.mutation.SetNote(s)
	return ouo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableNote(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetNote(*s)
	}
	return ouo
}

// ClearNote clears the value of the "note" field.
func (ouo *OrderUpdateOne) ClearNote() *OrderUpdateOne {
	ouo.mutation.ClearNote()
	return ouo
}

// SetTotal sets the "total" field.
func (ouo *OrderUpdateOne) SetTotal(f float32) *OrderUpdateOne {
	ouo.mutation.ResetTotal()
	ouo.mutation.SetTotal(f)
	return ouo
}

// AddTotal adds f to the "total" field.
func (ouo *OrderUpdateOne) AddTotal(f float32) *OrderUpdateOne {
	ouo.mutation.AddTotal(f)
	return ouo
}

// SetDiscount sets the "discount" field.
func (ouo *OrderUpdateOne) SetDiscount(f float32) *OrderUpdateOne {
	ouo.mutation.ResetDiscount()
	ouo.mutation.SetDiscount(f)
	return ouo
}

// AddDiscount adds f to the "discount" field.
func (ouo *OrderUpdateOne) AddDiscount(f float32) *OrderUpdateOne {
	ouo.mutation.AddDiscount(f)
	return ouo
}

// AddCartIDs adds the "cart" edge to the Cart entity by IDs.
func (ouo *OrderUpdateOne) AddCartIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.AddCartIDs(ids...)
	return ouo
}

// AddCart adds the "cart" edges to the Cart entity.
func (ouo *OrderUpdateOne) AddCart(c ...*Cart) *OrderUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ouo.AddCartIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OrderUpdateOne) SetUserID(id int) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetUserID(*id)
	}
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearCart clears all "cart" edges to the Cart entity.
func (ouo *OrderUpdateOne) ClearCart() *OrderUpdateOne {
	ouo.mutation.ClearCart()
	return ouo
}

// RemoveCartIDs removes the "cart" edge to Cart entities by IDs.
func (ouo *OrderUpdateOne) RemoveCartIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.RemoveCartIDs(ids...)
	return ouo
}

// RemoveCart removes "cart" edges to Cart entities.
func (ouo *OrderUpdateOne) RemoveCart(c ...*Cart) *OrderUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ouo.RemoveCartIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(order.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Phone(); ok {
		_spec.SetField(order.FieldPhone, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Note(); ok {
		_spec.SetField(order.FieldNote, field.TypeString, value)
	}
	if ouo.mutation.NoteCleared() {
		_spec.ClearField(order.FieldNote, field.TypeString)
	}
	if value, ok := ouo.mutation.Total(); ok {
		_spec.SetField(order.FieldTotal, field.TypeFloat32, value)
	}
	if value, ok := ouo.mutation.AddedTotal(); ok {
		_spec.AddField(order.FieldTotal, field.TypeFloat32, value)
	}
	if value, ok := ouo.mutation.Discount(); ok {
		_spec.SetField(order.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ouo.mutation.AddedDiscount(); ok {
		_spec.AddField(order.FieldDiscount, field.TypeFloat32, value)
	}
	if ouo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedCartIDs(); len(nodes) > 0 && !ouo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.CartTable,
			Columns: []string{order.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
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
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}