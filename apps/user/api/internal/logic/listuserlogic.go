package logic

import (
	"context"
	"go-zero-demo/apps/user/rpc/types/user"

	"go-zero-demo/apps/user/api/internal/svc"
	"go-zero-demo/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser() (resp []*types.ListUserResponse, err error) {
	result, err := l.svcCtx.UserRpc.ListUser(l.ctx, &user.ListUserRequest{})
	if err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return make([]*types.ListUserResponse, 0), nil
	}
	for _, item := range result.Data {
		resp = append(resp, &types.ListUserResponse{
			Id:     item.Id,
			Name:   item.Name,
			Gender: item.Gender,
			Mobile: item.Mobile,
		})
	}
	return resp, nil
}
