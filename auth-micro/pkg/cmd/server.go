package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/logger"
	"github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/protocol/grpc"
	"golang.org/x/net/context"
)

// Config is ...
type Config struct {
	GRPCPort            string
	DatastoreDBHost     string
	DatastoreDBUser     string
	DatastoreDBPassword string
	DatastoreDBSchema   string
	DatastoreDBPort     string
	LogLevel            int
	LogTimeFormat       string
}

// RunServer is ...
func RunServer() error {
	ctx := context.Background()

	fmt.Println("Hey")
	envErr := godotenv.Load()

	if envErr != nil {
		return fmt.Errorf("Error while loading .env file: %v", envErr)
	}

	logLevel, parseErr := strconv.ParseInt(os.Getenv("LOG_LEVEL"), 10, 64)

	if parseErr != nil {
		return fmt.Errorf("Error while parsing .env file: %v", parseErr)
	}

	config := Config{
		GRPCPort:            os.Getenv("PORT"),
		DatastoreDBHost:     os.Getenv("DB_HOST"),
		DatastoreDBUser:     os.Getenv("DB_USER"),
		DatastoreDBPassword: os.Getenv("DB_PASSWORD"),
		DatastoreDBPort:     os.Getenv("DB_PORT"),
		DatastoreDBSchema:   os.Getenv("DB_NAME"),
		LogLevel:            int(logLevel),
		LogTimeFormat:       os.Getenv("LOG_TIME_FORMAT"),
	}

	if err := logger.Init(config.LogLevel, config.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	param := "charset=utf8&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
		config.DatastoreDBUser, config.DatastoreDBPassword, config.DatastoreDBHost, config.DatastoreDBPort, config.DatastoreDBSchema, param)

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("Error while connection to database: %v", err)
	}

	fmt.Println("Connecting to the database...")
	defer db.Close()

	return grpc.RunServer(ctx, config.GRPCPort, db)
}
