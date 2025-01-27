package logic

import (
	"context"
	"go-zero-im/rpc/userclient"

	"go-zero-im/api/internal/svc"
	"go-zero-im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserRes, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.User.GetUser(l.ctx, &userclient.UserReq{
		Id: req.Id,
	})
	if err != nil {
		return resp, err
	}
	return &types.UserRes{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
