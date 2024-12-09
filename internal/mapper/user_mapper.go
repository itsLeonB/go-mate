package mapper

import (
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
)

func MapUserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func MapUsersToResponses(users []*entity.User) []*model.UserResponse {
	userResponses := make([]*model.UserResponse, len(users))
	for i := 0; i < len(users); i++ {
		userResponses[i] = MapUserToResponse(users[i])
	}

	return userResponses
}
