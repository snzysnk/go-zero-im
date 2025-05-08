package logic

import (
	"context"
	"strconv"

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
	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	userData, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	return &user.UserRes{
		Id:   strconv.Itoa(int(userData.Id)),
		Name: userData.Name,
	}, nil
}
