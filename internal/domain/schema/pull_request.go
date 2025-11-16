package schema

import (
	"AvitoTechTask/internal/domain/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// PullRequest holds the schema definition for the PullRequest entity.
type PullRequest struct {
	ent.Schema
}

// Fields of the PullRequest.
func (PullRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("pull_request_name").
			NotEmpty().Unique(),
		field.UUID("author_id", uuid.UUID{}),
		field.Int("status").
			GoType(types.PullRequestStatus(0)),
		field.JSON("assigned_reviewers", []uuid.UUID{}).
			Default([]uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("merged_at").
			Optional().
			Nillable(),
	}
}

// Edges of the PullRequest.
func (PullRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("authored_pull_requests").
			Field("author_id").
			Unique().
			Required(),
		edge.To("reviewers", User.Type),
	}
}

// Indexes of the PullRequest.
func (PullRequest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("author_id"),
		index.Fields("status"),
		index.Fields("created_at"),
	}
}
