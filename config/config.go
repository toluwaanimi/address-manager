package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go/build"
	"os"
	"path/filepath"
)

type Secrets struct {
	DatabaseURL            string `json:"DATABASE_URL"`
	Port                   string `json:"PORT"`
	Environment            Environment
	RedisURL               string `json:"REDIS_URL"`
	RedisUsername          string `json:"REDIS_USERNAME"`
	RedisPassword          string `json:"REDIS_PASSWORD"`
	ThresholdURL           string `json:"THRESHOLD_URL"`
	BTCWalletDepositID     string `json:"BTC_DEPOSIT_WALLET_SANDBOX_WALLET_ID"`
	BTCWalletDepositCode   string `json:"BTC_DEPOSIT_WALLET_SANDBOX_WALLET_TOKEN"`
	BTCWalletDepositSecret string `json:"BTC_DEPOSIT_WALLET_SANDBOX_WALLET_SECRET"`
}

var ss Secrets

const ServiceName = "ms-address"

func init() {
	importPath := fmt.Sprintf("%s/config", ServiceName)
	p, err := build.Default.Import(importPath, "", build.FindOnly)
	if err == nil {
		env := filepath.Join(p.Dir, "../.env")
		_ = godotenv.Load(env)
	}

	ss = Secrets{}

	ss.DatabaseURL = os.Getenv("PORT")
	ss.DatabaseURL = os.Getenv("DATABASE_URL")
	ss.RedisURL = os.Getenv("REDIS_URL")
	ss.RedisUsername = os.Getenv("REDIS_USERNAME")
	ss.RedisPassword = os.Getenv("REDIS_PASSWORD")
	ss.ThresholdURL = os.Getenv("THRESHOLD_URL")
	ss.BTCWalletDepositID = os.Getenv("BTC_DEPOSIT_WALLET_SANDBOX_WALLET_ID")
	ss.BTCWalletDepositCode = os.Getenv("BTC_DEPOSIT_WALLET_SANDBOX_WALLET_TOKEN")
	ss.BTCWalletDepositSecret = os.Getenv("BTC_DEPOSIT_WALLET_SANDBOX_WALLET_SECRET")

	ss.Environment = Environment(os.Getenv("ENV"))
	if err := ss.Environment.IsValid(); err != nil {
		logrus.Fatal("Error in environment variables: ", err)
	}

	if ss.Port = os.Getenv("PORT"); ss.Port == "" {
		ss.Port = ":8080"
	}
}

func GetSecrets() Secrets {
	return ss
}
