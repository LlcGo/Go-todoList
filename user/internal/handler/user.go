package handler

import (
	"context"
	"user/internal/logic"
	"user/internal/repository"
	"user/pkg/e"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *logic.UserRequest) (resp *logic.UserDetailResponse, err error) {
	var user repository.User
	resp = new(logic.UserDetailResponse)
	resp.Code = e.Success
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
}
