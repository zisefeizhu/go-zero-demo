package logic

import (
	"context"
	"go-zero-demo/apps/user/model"
	"go-zero-demo/apps/user/rpc/internal/svc"
	"go-zero-demo/apps/user/rpc/types/user"
	"go-zero-demo/pkg/cryptx"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	findOneByMobile, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, status.Error(100, err.Error())
	}

	if findOneByMobile != nil {
		return &user.RegisterResponse{
			Id:     findOneByMobile.Id,
			Name:   findOneByMobile.Name,
			Gender: findOneByMobile.Gender,
			Mobile: findOneByMobile.Mobile,
		}, nil
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		result, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = result.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}
	return nil, status.Error(500, err.Error())
}
