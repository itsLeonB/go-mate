package service

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/entity"
)

type scoringServiceNaive struct {
}

func NewScoringServiceNaive() ScoringService {
	return &scoringServiceNaive{}
}

func (ssn *scoringServiceNaive) ScoreAndSortUsers(ctx context.Context, users []*entity.User) []*entity.User {
	// Naive implementation of scoring and sorting users
	var scoredUsers []*entity.User

	for i := 0; i < len(users); i++ {
		if !users[i].DeletedAt.Valid || users[i].ID != ctx.Value(appconstant.ContextUserID) {
			scoredUsers = append(scoredUsers, users[i])
		}
	}

	if len(scoredUsers) < 10 {
		return scoredUsers
	}

	return scoredUsers[:10]
}
