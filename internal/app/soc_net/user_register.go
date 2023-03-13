package soc_net

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"otus/socNet/internal/hash"
	pb "otus/socNet/internal/pb"
	"otus/socNet/internal/structs"
)

func (i *Implementation) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	userData, err := convUserData(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	userID, err := i.socNetAPI.UserRegister(ctx, userData)
	return &pb.UserRegisterResponse{
		UserID: userID,
	}, nil
}

func convUserData(req *pb.UserRegisterRequest) (*structs.User, error) {
	if req.GetFirstName() == "" || req.GetSecondName() == "" {
		return nil, fmt.Errorf("empty name")
	}
	if req.GetCity() == "" {
		return nil, fmt.Errorf("empty city")
	}
	if req.GetPassword() == "" {
		return nil, fmt.Errorf("empty password")
	}
	newHash, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	return &structs.User{
		FirstName:  req.GetFirstName(),
		SecondName: req.GetSecondName(),
		City:       req.GetCity(),
		DateBirth:  time.Unix(req.GetDateBirth(), 0),
		Biography:  req.GetBiography(),
		Pass:       newHash,
	}, nil
}
