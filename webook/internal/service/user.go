package service

import (
	"context"
	"gin_test/webook/internal/domain"
	"gin_test/webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}
