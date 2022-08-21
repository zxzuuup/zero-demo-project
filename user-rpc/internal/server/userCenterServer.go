// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"zero-demo-study/user-rpc/internal/logic"
	"zero-demo-study/user-rpc/internal/svc"
	"zero-demo-study/user-rpc/pb"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserCenterServer) GetUserCreate(ctx context.Context, in *pb.GetUserCreateReq) (*pb.GetUserCreateResp, error) {
	l := logic.NewGetUserCreateLogic(ctx, s.svcCtx)
	return l.GetUserCreate(in)
}