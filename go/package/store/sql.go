package store

import (
	"database/sql"
	"go-python/internal/domain"
)

type Sql struct {
	DB *sql.DB 
}

func (s *Sql) GetTeamByRank(rank int) (*domain.Team, error) {
	var team domain.Team

	query := "SELECT * FROM regular_season WHERE rk = ?"

	row := s.DB.QueryRow(query, rank)
	err := row.Scan(&team.Unnamed, &team.Rank, &team.Squad, &team.MatchesPlayed, &team.Wins, &team.Draws, &team.Loses, &team.GoalsFavor, &team.GoalsAgainst, &team.GoalDifference, &team.Points, &team.PointsPerMatchPlayed, &team.ExpecGoals, &team.ExpecGoalsAllowed, &team.ExpecGoalsDifference, &team.ExpecGoalsDiffPerNinety, &team.Attendance, &team.TopTeamScorer, &team.Goalkeeper, &team.Notes)

	if err != nil {
		return nil, err
	}

	return &team, nil
}