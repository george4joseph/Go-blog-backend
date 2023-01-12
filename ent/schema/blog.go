package schema 

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"github.com/google/uuid"

	
)

type Blog struct {
	ent.Schema
}

func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Unique().Immutable(),
		field.String("Title"),
      	field.String("Content"),
      	field.Time("created_at").Default(time.Now),
		field.UUID("user_id", uuid.UUID{}).Unique().Optional(),
	}
}

func (Blog) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("owner",User.Type).
			Ref("blogs").
			Unique().
			Field("user_id"),
	}
}