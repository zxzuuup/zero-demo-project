package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo-study/user-rpc/internal/config"
	"zero-demo-study/user-rpc/model"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserDataModel model.UserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//根据config中的的连接信息，生从对应的model对象
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
		UserDataModel: model.NewUserDataModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
	}
}
