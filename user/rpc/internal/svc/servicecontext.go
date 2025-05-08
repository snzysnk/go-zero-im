package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-im/models"
	"go-zero-im/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserModel(mysql, c.Cache),
	}
}
