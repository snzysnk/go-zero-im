package logic

import (
	"context"

	"go-zero-im/rpc/internal/svc"
	"go-zero-im/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.UserReq) (*user.UserRes, error) {
	return &user.UserRes{
		Id:   in.Id,
		Name: "test",
	}, nil
}
