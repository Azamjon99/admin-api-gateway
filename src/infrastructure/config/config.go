package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// debug, info, warn, error,
	LogLevel         string
	HttpPort         string
	GrpcPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	SmsProvideApiKey string
	JWTSecret        string
	JWTExpiresInSec  int
	LogisticStaffAddress        string
	LogisticsVehicleService     string
}

func Load() (Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	var config Config
	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("HTTP_PORT", ":1000")
	v.SetDefault("LOGISTICS_STAFF_ADDRESS", "localhost:3535")
	v.SetDefault("LOGISTIC_VEHICLE_ADDRESS", "localhost:4242")
	v.SetDefault("JWT_SECRET", "")
	v.SetDefault("JWT_EXPIRES_IN_SEC", 2630000) // 1 month


	config.LogLevel = v.GetString("LOG_LEVEL")
	config.HttpPort = v.GetString("HTTP_PORT")
	config.LogisticStaffAddress = v.GetString("LOGISTICS_STAFF_ADDRESS")
	config.LogisticsVehicleService = v.GetString("LOGISTIC_VEHICLE_ADDRESS")
	config.JWTSecret = v.GetString("JWT_SECRET")
	config.JWTExpiresInSec = v.GetInt("JWT_EXPIRES_IN_SEC")

	return config, nil
}

func (c Config) NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	level, err := zapcore.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}

	config.Level.SetLevel(level)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
