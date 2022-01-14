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
	DatabaseURL string `json:"DATABASE_URL"`
	Port        string `json:"PORT"`
	Environment Environment
}

var ss Secrets

const ServiceName = "address.generator"

func initConfig() {
	importPath := fmt.Sprintf("%s/config", ServiceName)
	p, err := build.Default.Import(importPath, "", build.FindOnly)
	if err == nil {
		env := filepath.Join(p.Dir, "../.env")
		_ = godotenv.Load(env)
	}

	ss = Secrets{}

	ss.DatabaseURL = os.Getenv("PORT")
	ss.DatabaseURL = os.Getenv("DATABASE_URL")

	ss.Environment = Environment(os.Getenv("ENV"))
	if err := ss.Environment.IsValid(); err != nil {
		logrus.Fatal("Error in environment variables: ", err)
	}
}

func GetSecrets() Secrets {
	return ss
}
