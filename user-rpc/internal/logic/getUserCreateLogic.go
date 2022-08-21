package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zero-demo-study/user-rpc/model"

	"zero-demo-study/user-rpc/internal/svc"
	"zero-demo-study/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCreateLogic {
	return &GetUserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCreateLogic) GetUserCreate(in *pb.GetUserCreateReq) (*pb.GetUserCreateResp, error) {
	if err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{}
		user.Mobile = in.Mobile
		user.Nickname = in.Nickname
		//添加user
		dbResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		userId, _ := dbResult.LastInsertId()
		//添加userData
		userData := &model.UserData{}
		userData.UserId = userId
		userData.Data = in.Data
		userData.CreateTime = time.Now()
		userData.UpdateTime = time.Now()
		logx.Info(userData)
		if _, err := l.svcCtx.UserDataModel.TransInsert(ctx, session, userData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}
	return &pb.GetUserCreateResp{
		Flag: true,
	}, nil
}
