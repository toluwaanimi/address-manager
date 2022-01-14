package config

import (
	"fmt"
	"ms-address/log"
	"strings"
)

type Environment string

const (
	Local      Environment = "local"
	Staging    Environment = "staging"
	Production Environment = "production"
)

func (e Environment) IsValid() error {
	switch e {
	case Local, Staging, Production:
		return nil
	default:
		return fmt.Errorf("unknown environment : %s", e)
	}
}

func (e Environment) String() string { return strings.ToLower(string(e)) }

func (e Environment) Is(env Environment) bool {
	return e == env
}

func (e Environment) LogLevel() log.Level {
	switch e {
	case Staging, Production:
		return log.LevelInfo
	case Local:
		fallthrough
	default:
		return log.LevelDebug
	}
}
