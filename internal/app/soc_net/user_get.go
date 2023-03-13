package soc_net

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "otus/socNet/internal/pb"
	"otus/socNet/internal/structs"
)

func (i *Implementation) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID, err := convUserID(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	statusGetting := structs.StatusUserDBNotValid
	if userID >= 0 {
		statusGetting, err = i.socNetAPI.GetUser(ctx, userID)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

	}

	return &pb.GetUserResponse{
		Status: statusGetting.GetDescription(),
	}, nil
}

func convUserID(req *pb.GetUserRequest) (int64, error) {
	id, err := strconv.ParseInt(req.GetUserID(), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
