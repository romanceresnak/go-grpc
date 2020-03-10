package users

import (
	"context"
	pb "github.com/romanceresnak/go-grpc/pb"
)

type grpcHandler struct {
}

func GetRoutes() pb.V1UsersServer {
	return &grpcHandler{}
}

func (h *grpcHandler) Create(ctx context.Context, req *pb.CreateUserRequest) (res *pb.UserReply, err error) {
	res = new(pb.UserReply)

	return res, nil
}

func (h *grpcHandler) FindById(ctx context.Context, req *pb.FindByIdRequest) (res *pb.UserReply, err error) {
	res = new(pb.UserReply)

	return res, nil
}

func (h *grpcHandler) FindByEmail(ctx context.Context, req *pb.FindByEmailRequest) (res *pb.UserReply, err error) {
	res = new(pb.UserReply)

	return res, nil
}

func (h *grpcHandler) Update(ctx context.Context, req *pb.UpdateUserRequest) (res *pb.UserReply, err error) {
	res = new(pb.UserReply)

	return res, nil
}
