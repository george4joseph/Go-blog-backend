package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Unique().Immutable(),
		field.String("name").MaxLen(250),
		field.Time("created_at").Default(time.Now),
		field.String("email").Match(regexp.MustCompile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)).Unique(),
		field.Enum("user_type").GoType(UserType("")),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blogs", Blog.Type),
	}
}

type ( 
	UserType string
)

const (
	User_ UserType = "USER"
	Admin UserType = "ADMIN"
)

func (UserType) Values() (kinds []string) {
	for _, s := range []UserType{User_, Admin} {
		kinds = append(kinds, string(s))
	}

	return
}
