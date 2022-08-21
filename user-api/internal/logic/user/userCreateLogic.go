package user

import (
	"context"
	"zero-demo-study/user-rpc/pb"

	"zero-demo-study/user-api/internal/svc"
	"zero-demo-study/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserCreate(l.ctx, &pb.GetUserCreateReq{
		Nickname: req.NickName,
		Mobile:   req.Mobile,
		Data:     req.Data,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserCreateResp{
		Flag: userResp.Flag,
	}, nil
}
