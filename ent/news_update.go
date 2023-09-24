// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nghuyentrangvt98/banhmioi/ent/news"
	"github.com/nghuyentrangvt98/banhmioi/ent/predicate"
)

// NewsUpdate is the builder for updating News entities.
type NewsUpdate struct {
	config
	hooks    []Hook
	mutation *NewsMutation
}

// Where appends a list predicates to the NewsUpdate builder.
func (nu *NewsUpdate) Where(ps ...predicate.News) *NewsUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTitle sets the "title" field.
func (nu *NewsUpdate) SetTitle(s string) *NewsUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetSubtitle sets the "subtitle" field.
func (nu *NewsUpdate) SetSubtitle(s string) *NewsUpdate {
	nu.mutation.SetSubtitle(s)
	return nu
}

// SetContent sets the "content" field.
func (nu *NewsUpdate) SetContent(s string) *NewsUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetAuthor sets the "author" field.
func (nu *NewsUpdate) SetAuthor(s string) *NewsUpdate {
	nu.mutation.SetAuthor(s)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NewsUpdate) SetCreatedAt(t time.Time) *NewsUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetCategory sets the "category" field.
func (nu *NewsUpdate) SetCategory(s string) *NewsUpdate {
	nu.mutation.SetCategory(s)
	return nu
}

// SetProductURL sets the "product_url" field.
func (nu *NewsUpdate) SetProductURL(s string) *NewsUpdate {
	nu.mutation.SetProductURL(s)
	return nu
}

// SetImageURL sets the "image_url" field.
func (nu *NewsUpdate) SetImageURL(s string) *NewsUpdate {
	nu.mutation.SetImageURL(s)
	return nu
}

// Mutation returns the NewsMutation object of the builder.
func (nu *NewsUpdate) Mutation() *NewsMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NewsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NewsUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NewsUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NewsUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NewsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(news.Table, news.Columns, sqlgraph.NewFieldSpec(news.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Subtitle(); ok {
		_spec.SetField(news.FieldSubtitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
	}
	if value, ok := nu.mutation.Author(); ok {
		_spec.SetField(news.FieldAuthor, field.TypeString, value)
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(news.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Category(); ok {
		_spec.SetField(news.FieldCategory, field.TypeString, value)
	}
	if value, ok := nu.mutation.ProductURL(); ok {
		_spec.SetField(news.FieldProductURL, field.TypeString, value)
	}
	if value, ok := nu.mutation.ImageURL(); ok {
		_spec.SetField(news.FieldImageURL, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NewsUpdateOne is the builder for updating a single News entity.
type NewsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsMutation
}

// SetTitle sets the "title" field.
func (nuo *NewsUpdateOne) SetTitle(s string) *NewsUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetSubtitle sets the "subtitle" field.
func (nuo *NewsUpdateOne) SetSubtitle(s string) *NewsUpdateOne {
	nuo.mutation.SetSubtitle(s)
	return nuo
}

// SetContent sets the "content" field.
func (nuo *NewsUpdateOne) SetContent(s string) *NewsUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetAuthor sets the "author" field.
func (nuo *NewsUpdateOne) SetAuthor(s string) *NewsUpdateOne {
	nuo.mutation.SetAuthor(s)
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NewsUpdateOne) SetCreatedAt(t time.Time) *NewsUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetCategory sets the "category" field.
func (nuo *NewsUpdateOne) SetCategory(s string) *NewsUpdateOne {
	nuo.mutation.SetCategory(s)
	return nuo
}

// SetProductURL sets the "product_url" field.
func (nuo *NewsUpdateOne) SetProductURL(s string) *NewsUpdateOne {
	nuo.mutation.SetProductURL(s)
	return nuo
}

// SetImageURL sets the "image_url" field.
func (nuo *NewsUpdateOne) SetImageURL(s string) *NewsUpdateOne {
	nuo.mutation.SetImageURL(s)
	return nuo
}

// Mutation returns the NewsMutation object of the builder.
func (nuo *NewsUpdateOne) Mutation() *NewsMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NewsUpdate builder.
func (nuo *NewsUpdateOne) Where(ps ...predicate.News) *NewsUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NewsUpdateOne) Select(field string, fields ...string) *NewsUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated News entity.
func (nuo *NewsUpdateOne) Save(ctx context.Context) (*News, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NewsUpdateOne) SaveX(ctx context.Context) *News {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NewsUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NewsUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NewsUpdateOne) sqlSave(ctx context.Context) (_node *News, err error) {
	_spec := sqlgraph.NewUpdateSpec(news.Table, news.Columns, sqlgraph.NewFieldSpec(news.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "News.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, news.FieldID)
		for _, f := range fields {
			if !news.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != news.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Subtitle(); ok {
		_spec.SetField(news.FieldSubtitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Author(); ok {
		_spec.SetField(news.FieldAuthor, field.TypeString, value)
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(news.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Category(); ok {
		_spec.SetField(news.FieldCategory, field.TypeString, value)
	}
	if value, ok := nuo.mutation.ProductURL(); ok {
		_spec.SetField(news.FieldProductURL, field.TypeString, value)
	}
	if value, ok := nuo.mutation.ImageURL(); ok {
		_spec.SetField(news.FieldImageURL, field.TypeString, value)
	}
	_node = &News{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
