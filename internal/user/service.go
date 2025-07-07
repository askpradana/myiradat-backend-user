package user

import (
	"context"
	"errors"
	"iradat/profile/internal/user/dto"
)

type InterfaceUserService interface {
	GetUserDetail(ctx context.Context, detail dto.RequestGetDetail) (res dto.ResponseGetDetail, err error)
	ListUser(ctx context.Context, detail dto.ResponseGetListUser) (res dto.ResponseGetListUser, err error)
}

type UserService struct {
}

func (u *UserService) ListUser(ctx context.Context, detail dto.ResponseGetListUser) (res dto.ResponseGetListUser, err error) {
	return dto.ResponseGetListUser{}, nil
}

func (u *UserService) GetUserDetail(ctx context.Context, detail dto.RequestGetDetail) (res dto.ResponseGetDetail, err error) {
	res = dto.ResponseGetDetail{
		UserID:   "",
		Username: "",
	}

	if res.UserID == "" {
		msg := "UserID is empty"
		return res, errors.New(msg)
	}

	return res, nil
}

func NewUserService() InterfaceUserService {
	return &UserService{}
}
