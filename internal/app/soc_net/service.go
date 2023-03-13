package soc_net

import (
	pb "otus/socNet/internal/pb"
)

type Implementation struct {
	socNetAPI SocNetAPI
	pb.UnimplementedSocNetServer
}

// NewSocNetAPI return new instance of Implementation.
func NewSocNetAPI(socNetAPI SocNetAPI) *Implementation {
	return &Implementation{socNetAPI: socNetAPI}
}
