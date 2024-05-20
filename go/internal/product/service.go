package product

import "go-python/internal/domain"

type IService interface {
	GetTeamByRank(rank int) (*domain.Team, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetTeamByRank(rank int) (*domain.Team, error) {
	team, err := s.Repository.GetTeamByRank(rank)

	if err != nil {
		return nil, err
	}

	return team, nil
}