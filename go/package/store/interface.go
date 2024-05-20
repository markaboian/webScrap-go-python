package store

import "go-python/internal/domain"

type Interface interface {
	GetTeamByRank(rank int) (*domain.Team, error)
}