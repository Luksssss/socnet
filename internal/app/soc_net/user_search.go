package soc_net

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "otus/socNet/internal/pb/api/socnet"
	"otus/socNet/internal/structs"
)

func (i *Implementation) UserSearch(ctx context.Context, req *pb.UserSearchRequest) (*pb.UserSearchResponse, error) {
	usersSearch, err := convUserSearch(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	usersSearchRes, err := i.socNetAPI.UserSearch(ctx, usersSearch)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := make([]*pb.UserInfo, 0, len(usersSearchRes))
	for _, user := range usersSearchRes {
		res = append(res, &pb.UserInfo{
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			DateBirth:  user.DateBirth.String(),
		})
	}

	return &pb.UserSearchResponse{
		UserInfo: res,
	}, nil
}

func convUserSearch(req *pb.UserSearchRequest) (*structs.UserSearch, error) {
	if req.GetFirstName() == "" || req.GetSecondName() == "" {
		return nil, fmt.Errorf("empty firstName or SecondName")
	}
	return &structs.UserSearch{
		FirstName:  strings.ToLower(req.GetFirstName()),
		SecondName: strings.ToLower(req.GetSecondName()),
	}, nil
}
