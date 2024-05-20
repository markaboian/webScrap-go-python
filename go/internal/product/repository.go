package product

import (
	"go-python/internal/domain"
	"go-python/package/store"
)

type IRepository interface {
	GetTeamByRank(rank int) (*domain.Team, error)
}

type Repository struct {
	Interface store.Interface
}

func (r *Repository) GetTeamByRank(rank int) (*domain.Team, error) {
	team, err := r.Interface.GetTeamByRank(rank)
	if err != nil {
		return nil, err
	}

	return team, nil
}