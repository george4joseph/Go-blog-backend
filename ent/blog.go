// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/george4joseph/go-blog-backend/ent/blog"
	"github.com/george4joseph/go-blog-backend/ent/user"
	"github.com/google/uuid"
)

// Blog is the model entity for the Blog schema.
type Blog struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "Title" field.
	Title string `json:"Title,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlogQuery when eager-loading is set.
	Edges BlogEdges `json:"edges"`
}

// BlogEdges holds the relations/edges for other nodes in the graph.
type BlogEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlogEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blog.FieldTitle, blog.FieldContent:
			values[i] = new(sql.NullString)
		case blog.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case blog.FieldID, blog.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blog fields.
func (b *Blog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case blog.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case blog.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Content", values[i])
			} else if value.Valid {
				b.Content = value.String
			}
		case blog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case blog.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				b.UserID = *value
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Blog entity.
func (b *Blog) QueryOwner() *UserQuery {
	return (&BlogClient{config: b.config}).QueryOwner(b)
}

// Update returns a builder for updating this Blog.
// Note that you need to call Blog.Unwrap() before calling this method if this Blog
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blog) Update() *BlogUpdateOne {
	return (&BlogClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blog) Unwrap() *Blog {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blog is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blog) String() string {
	var builder strings.Builder
	builder.WriteString("Blog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("Title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("Content=")
	builder.WriteString(b.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", b.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Blogs is a parsable slice of Blog.
type Blogs []*Blog

func (b Blogs) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
