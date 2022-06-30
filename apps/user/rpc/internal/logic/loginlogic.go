package logic

import (
	"context"
	"go-zero-demo/apps/user/model"
	"go-zero-demo/pkg/cryptx"
	"google.golang.org/grpc/status"

	"go-zero-demo/apps/user/rpc/internal/svc"
	"go-zero-demo/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
   查询用户是否存在
   判断密码是否正确
*/

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	result, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	passwordEncrypt := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if passwordEncrypt != result.Password {
		return nil, status.Error(100, "密码错误")
	}
	return &user.LoginResponse{
		Id:     result.Id,
		Name:   result.Name,
		Gender: result.Gender,
		Mobile: result.Mobile,
	}, nil

}
