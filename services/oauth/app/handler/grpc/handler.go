package grpc

import (
	"demo04/services/oauth/app/bootstrap"
	"demo04/services/oauth/app/repository"
	"demo04/services/oauth/protobuf/pb"
)

type Handler struct {
	pb.UnimplementedOauthServiceServer
	repository *repository.Repository
}

// New :
func New(bs *bootstrap.Bootstrap) *Handler {
	h := new(Handler)
	h.repository = bs.Repository
	return h
}
