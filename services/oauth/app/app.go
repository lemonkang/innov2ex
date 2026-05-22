package app

import (
	"context"
	"demo04/services/oauth/app/bootstrap"
	gh "demo04/services/oauth/app/handler/grpc"
	"demo04/services/oauth/env"
	"demo04/services/oauth/protobuf/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedOauthServiceServer
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		UserKey:     "mocked_token",
		AccessToken: "mocked_token",
	}, nil
}

// 拦截器1：日志拦截器
func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		log.Printf("请求开始：%s, 参数：%v", info.FullMethod, req)
		// 执行后续拦截器/业务逻辑
		resp, err = handler(ctx, req)
		log.Printf("请求结束：%s, 错误：%v", info.FullMethod, err)
		return resp, err
	}
}

// 拦截器2：鉴权拦截器
func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// 模拟鉴权失败
		if false {
			return nil, status.Errorf(codes.Unauthenticated, "未授权")
		}
		// 鉴权通过，执行后续逻辑
		return handler(ctx, req)
	}
}
func Start(grpcPort, httpPort string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	env.Init()
	bs := bootstrap.New(ctx)
	defer bs.Release()
	h := gh.New(bs)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			LogInterceptor(),  // 第一个执行
			AuthInterceptor(), // 第二个执行
		),
	)

	pb.RegisterOauthServiceServer(grpcServer, h)

	log.Println("grpc server start :50051")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
