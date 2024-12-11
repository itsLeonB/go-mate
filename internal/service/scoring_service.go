package service

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/util"
)

type scoringServiceNaive struct {
}

func NewScoringServiceNaive() ScoringService {
	return &scoringServiceNaive{}
}

func (ssn *scoringServiceNaive) ScoreAndSortUsers(ctx context.Context, users []*entity.User, isExtraRecommendation bool) ([]*entity.User, error) {
	// Naive implementation of scoring and sorting users
	userID, err := util.GetUUIDFromContext(ctx, appconstant.ContextUserID)
	if err != nil {
		return nil, err
	}

	var scoredUsers []*entity.User

	for i := 0; i < len(users); i++ {
		if !users[i].DeletedAt.Valid && users[i].ID != userID {
			scoredUsers = append(scoredUsers, users[i])
		}
	}

	maxUserCount := appconstant.MaxUserCountPerDay
	if isExtraRecommendation {
		maxUserCount *= appconstant.MaxUserCountExtraMultiplier
	}

	if len(scoredUsers) < maxUserCount {
		return scoredUsers, nil
	}

	return scoredUsers[:maxUserCount], nil
}
