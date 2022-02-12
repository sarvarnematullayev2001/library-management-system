package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDB       string
	LogLevel         string

	UserServiceHost string
	UserServicePort int

	ProducerServiceHost string
	ProducerServicePort int

	RPCPort string

	PasscodePool   string
	PasscodeLength int
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "library"))
	c.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "test1234"))

	c.UserServiceHost = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(getOrReturnDefaultValue("USER_SERVICE_PORT", 5002))

	c.ProducerServiceHost = cast.ToString(getOrReturnDefaultValue("PRODUCER_SERVICE_HOST", "localhost"))
	c.ProducerServicePort = cast.ToInt(getOrReturnDefaultValue("PRODUCER_SERVICE_PORT", 5003))

	c.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefaultValue("RPC_PORT", ":5005"))

	c.PasscodePool = cast.ToString(getOrReturnDefaultValue("PASSCODE_POOL", "0123456789"))
	c.PasscodeLength = cast.ToInt(getOrReturnDefaultValue("PASSCODE_LENGTH", "6"))

	return c
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
