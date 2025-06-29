package user

import "context"

type RequestGetDetail struct {
}

type ResponseGetDetail struct {
}

type InterfaceUserService interface {
	GetUserDetail(ctx context.Context, detail RequestGetDetail) (res ResponseGetDetail, err error)
}

type UserService struct {
}

func (u UserService) GetUserDetail(ctx context.Context, detail RequestGetDetail) (res ResponseGetDetail, err error) {
	res = ResponseGetDetail{}
	return res, nil
}

func NewUserService() InterfaceUserService {
	return &UserService{}
}
