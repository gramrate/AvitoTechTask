package team

import (
	"AvitoTechTask/pkg/ent"
)

type Repo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) *Repo {
	return &Repo{client: client}
}
