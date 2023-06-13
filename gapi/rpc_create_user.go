package gapi

import (
	"context"

	db "github.com/angrypenguin1995/simple__bank/db/sqlc"
	"github.com/angrypenguin1995/simple__bank/pb"
	"github.com/angrypenguin1995/simple__bank/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password :%s", err)
	}

	args := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "User %s already exists", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user : %s", err)
	}
	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return resp, nil
}
