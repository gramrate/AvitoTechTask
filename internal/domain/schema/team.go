// internal/ent/schema/team.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),

		field.String("team_name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", User.Type),
	}
}

// Indexes of the Team.
func (Team) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("team_name"),
	}
}
