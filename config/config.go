package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	TimeExpiredAt = time.Hour * 24
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost           string
	PostgresUser           string
	PostgresDatabase       string
	PostgresPassword       string
	PostgresPort           string
	PostgresMaxConnections int32

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	DefaultOffset int
	DefaultLimit  int

	SecretKey string
}

func Load() Config {

	if err := godotenv.Load("./app.env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	cfg.ServerPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":4001"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "khumoyun"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "admin@111"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "bikestore"))
	cfg.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAXCONS", 20))

	cfg.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost:"))
	cfg.RedisPort = cast.ToString(getOrReturnDefaultValue("REDIS_PORT", 6379))
	cfg.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", ""))
	cfg.RedisDB = cast.ToInt(getOrReturnDefaultValue("REDIS_DB", "khumoyun"))

	cfg.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "secretkey"))

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
