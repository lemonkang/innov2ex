package grpc

import (
	"context"
	"demo04/services/oauth/protobuf/pb"
	"fmt"

	"github.com/RevenueMonster/sqlike/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	masterUser, err := h.repository.FindMasterUserByEmail(ctx, req.GetEmail())
	fmt.Printf("查询的MasterUser: %v\n", masterUser)
	fmt.Printf("查询的MasterUser密码哈希: %v\n", string(masterUser.PasswordHash))
	if err != nil {
		return nil, err
	}
	if !masterUser.VerifyPassword(req.GetPassword()) {
		return nil, status.Error(codes.PermissionDenied, "invalid password")
	}
	reqUserKey, err := types.DecodeKey(req.UserKey)
	fmt.Printf("请求中的UserKey: %v\n", reqUserKey.Root().ID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user key")
	}

	// 模拟登录成功，返回用户信息和访问令牌
	return &pb.LoginResponse{
		UserKey:     masterUser.Key.String(),
		AccessToken: "mocked_access_token",
	}, nil
}
