package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-im/api/internal/svc"
	"go-zero-im/api/internal/types"
	"go-zero-im/rpc/userclient"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateReq) (resp *types.CreateRes, err error) {
	createUser, err := l.svcCtx.User.CreateUser(l.ctx, &userclient.CreateReq{
		Name: req.Name,
	})
	if err != nil {
		return resp, nil
	}
	return &types.CreateRes{
		Id: createUser.Id,
	}, nil
}
