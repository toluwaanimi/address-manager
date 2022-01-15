package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"ms-address/core"
	"ms-address/log"
	address "ms-address/pb"
	"net"
)

func ExecGRPCServer(portAddress string, core *core.AddressService, logger log.Logger) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Recovered from err: %v", err)
		}
	}()
	api := NewApiService(core, logger)
	server := grpc.NewServer()
	address.RegisterAddressServiceServer(server, api)

	lis, err := net.Listen("tcp", portAddress)
	if err != nil {
		logrus.Fatalf("failed to listen with the following errors: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}
