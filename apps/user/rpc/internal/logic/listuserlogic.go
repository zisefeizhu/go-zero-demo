package logic

import (
	"context"
	"fmt"
	"go-zero-demo/apps/user/model"
	"google.golang.org/grpc/status"

	"go-zero-demo/apps/user/rpc/internal/svc"
	"go-zero-demo/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *user.ListUserRequest) (*user.ListUserResponse, error) {
	results, err := l.svcCtx.UserModel.FindList()
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	resp := make([]*user.UserInfoResponse, 0, len(results))

	for _, item := range results {
		resp = append(resp, &user.UserInfoResponse{
			Id:     item.Id,
			Name:   item.Name,
			Gender: item.Gender,
			Mobile: item.Mobile,
		})
	}
	fmt.Println(resp)
	return &user.ListUserResponse{
		Data: resp,
	}, nil
}
