package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/koropati/go-portfolio/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(c context.Context, user domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Create(ctx, user)
}

func (u *userUsecase) Retrieve(c context.Context, filter domain.Filter) (users []domain.User, meta domain.MetaResponse, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Retrieve(ctx, filter)
}

func (u *userUsecase) Update(c context.Context, id uuid.UUID, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Update(ctx, id, user)
}

func (u *userUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Delete(ctx, id)
}
