package soc_net

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"otus/socNet/internal/hash"
	pb "otus/socNet/internal/pb/api/socnet"
	"otus/socNet/internal/structs"
)

func (i *Implementation) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userLogin, err := convLogin(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	statusLogin := structs.StatusLoginNotValid
	if userLogin != nil && userLogin.ID > 0 {
		userLogin, err = i.socNetAPI.Login(ctx, userLogin)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		statusLogin = structs.StatusLoginUserNotFound
		if userLogin.Hash != "" {
			statusLogin = structs.StatusLoginNotValid
			if hash.CheckPasswordHash(userLogin.Pass, userLogin.Hash) {
				statusLogin = structs.StatusLoginOK
			}
		}
	}

	return &pb.LoginResponse{
		Status: statusLogin.GetDescription(),
	}, nil
}

func convLogin(req *pb.LoginRequest) (*structs.UserLogin, error) {
	if req.GetUserID() == "" {
		return nil, fmt.Errorf("empty userID")
	}
	id, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	if req.GetPassword() == "" {
		return nil, fmt.Errorf("empty password")
	}

	return &structs.UserLogin{
		ID:   id,
		Pass: req.Password,
	}, nil
}
