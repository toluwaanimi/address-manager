package server

import (
	"context"
	"ms-address/core"
	"ms-address/log"
	address "ms-address/pb"
)

type grpcService struct {
	core   *core.AddressService
	logger log.Logger
}

func (grpc *grpcService) GenerateDepositAddress(ctx context.Context, request *address.GenerateDepositAddressRequest) (*address.GenerateDepositAddressResponse, error) {
	log.Info("&s", request)
	reply := address.GenerateDepositAddressResponse{
		Address: "address",
	}
	return &reply, nil
}

func NewApiService(core *core.AddressService, logger log.Logger) address.AddressServiceServer {
	return &grpcService{
		core:   core,
		logger: logger,
	}
}
