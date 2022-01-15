package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"ms-address/cache"
	"ms-address/config"
	"ms-address/core"
	"ms-address/log"
	"time"

	//"ms-address/config"
	//"ms-address/log"
	"ms-address/pkg/db"
	"ms-address/server"
)

const CacheKeyPrefix = "ACCOUNTS_DAILY_TRANSACTION_LIMIT::"

func main() {
	secrets := config.GetSecrets()
	logger := log.New(log.GetLevel(), 0)
	redisConn := redis.NewClient(&redis.Options{
		Addr:     secrets.RedisURL,
		Username: secrets.RedisUsername,
		Password: secrets.RedisPassword,
	})
	kvStore := cache.NewCache(redisConn, time.Hour*24, CacheKeyPrefix)
	service := core.NewAccountService(kvStore, logger)
	db.Init(secrets.DatabaseURL)
	//	h := handlers.New(DB)
	address := fmt.Sprintf("0.0.0.0:%s", secrets.Port)
	fmt.Printf("[*] %s listening on address: %s", config.ServiceName, address)
	server.ExecGRPCServer(secrets.Port, service, logger)
}
