package user

import (
	"context"
	"zero-demo-study/user-rpc/pb"

	"zero-demo-study/user-api/internal/svc"
	"zero-demo-study/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		UserId:   userResp.Id,
		NickName: userResp.Nickname,
		Mobile:   userResp.Mobile,
	}, nil
}
