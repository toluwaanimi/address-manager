package server

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-address/config"
	"ms-address/core"
	"ms-address/log"
	address "ms-address/pb"
	"ms-address/pkg/threshold"
)

type grpcService struct {
	core   *core.AddressService
	logger log.Logger
}

func (grpc *grpcService) GenerateDepositAddress(ctx context.Context, request *address.GenerateDepositAddressRequest) (*address.GenerateDepositAddressResponse, error) {
	log.Info("&s", request)
	walletConfig, _ := config.GetThresholdDepositWallet("BTC")
	resp, err := threshold.MakeRequest("GET", fmt.Sprintf("/v1/sofa/wallets/%s/addresses", walletConfig.WalletId),
		nil, nil)
	if err != nil {
		log.Error("GetDepositWalletAddresses failed", err)
	}
	var m map[string]interface{}
	e := json.Unmarshal(resp, &m)
	if e != nil {
		return nil, e
	}
	fmt.Printf("%+v\n", m)
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
