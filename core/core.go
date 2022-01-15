package core

import (
	"ms-address/cache"
	"ms-address/log"
)

type AddressService struct {
	kvStore cache.Cache
	logger  log.Logger
}

func NewAccountService(
	kvStore cache.Cache,
	logger log.Logger,
) *AddressService {
	return &AddressService{
		kvStore: kvStore,
		logger:  logger,
	}
}
