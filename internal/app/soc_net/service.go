package soc_net

import (
	pb "otus/socNet/internal/pb/api/socnet"
)

type Implementation struct {
	socNetAPI SocNetAPI
	pb.UnimplementedSocNetServer
}

// NewSocNetAPI return new instance of Implementation.
func NewSocNetAPI(socNetAPI SocNetAPI) *Implementation {
	return &Implementation{socNetAPI: socNetAPI}
}
