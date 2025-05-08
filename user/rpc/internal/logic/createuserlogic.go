package logic

import (
	"context"
	"go-zero-im/models"
	"strconv"

	"go-zero-im/rpc/internal/svc"
	"go-zero-im/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 添加 CreateUser 来测试在 go-zero 中使用mysql
func (l *CreateUserLogic) CreateUser(in *user.CreateReq) (*user.CreateRes, error) {
	// todo: add your logic here and delete this line
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	id, _ := insert.LastInsertId()

	return &user.CreateRes{
		Id: strconv.Itoa(int(id)),
	}, nil
}
