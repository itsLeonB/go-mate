package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/apperror"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/repository"
	"github.com/itsLeonB/go-mate/internal/util"
	"github.com/rotisserie/eris"
)

type authServiceImpl struct {
	userRepository repository.UserRepository
	hasher         util.Hash
	jwt            util.JWT
}

func NewAuthService(userRepository repository.UserRepository, hasher util.Hash, jwt util.JWT) AuthService {
	return &authServiceImpl{
		userRepository: userRepository,
		hasher:         hasher,
		jwt:            jwt,
	}
}

func (as *authServiceImpl) Register(ctx context.Context, request *model.RegisterRequest) error {
	user, err := as.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return err
	}
	if user != nil {
		return apperror.DuplicateEmailError(request.Email)
	}

	hash, err := as.hasher.Hash(request.Password)
	if err != nil {
		return eris.Wrap(err, "error hashing password")
	}

	user = &entity.User{
		Email:    request.Email,
		Password: hash,
	}

	err = as.userRepository.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (as *authServiceImpl) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := as.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperror.LoginError()
	}

	ok, err := as.hasher.CheckHash(user.Password, request.Password)
	if err != nil {
		return nil, eris.Wrap(err, "error checking password hash")
	}
	if !ok {
		return nil, apperror.LoginError()
	}

	token, err := as.jwt.CreateToken(gin.H{appconstant.ContextUserID: user.ID})
	if err != nil {
		return nil, eris.Wrap(err, "error creating token")
	}

	return model.NewLoginResponse(token), nil
}
