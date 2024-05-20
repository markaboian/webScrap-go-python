package domain

import "database/sql"

type Team struct {
	Unnamed                 int            `json:"unnamed"`
	Rank                    int            `json:"rk"`
	Squad                   string         `json:"squad"`
	MatchesPlayed           int            `json:"mp"`
	Wins                    int            `json:"wins"`
	Draws                   int            `json:"draws"`
	Loses                   int            `json:"loses"`
	GoalsFavor              int            `json:"goalsFavor"`
	GoalsAgainst            int            `json:"goalsAgainst"`
	GoalDifference          int            `json:"goalDiference"`
	Points                  int            `json:"points"`
	PointsPerMatchPlayed    float64        `json:"pointsPerMatchPlayed"`
	ExpecGoals              float64        `json:"expectedGoals"`
	ExpecGoalsAllowed       float64        `json:"expectedGoalsAllowed"`
	ExpecGoalsDifference    float64        `json:"expectedGoalsDifference"`
	ExpecGoalsDiffPerNinety float64        `json:"expectedGoalsDiffPerNinety"`
	Attendance              int            `json:"attendance"`
	TopTeamScorer           string         `json:"topTeamScorer"`
	Goalkeeper              string         `json:"goalkeeper"`
	Notes                   sql.NullString `json:"notes"`
}
