package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL   string
	MongoDB    string
	ServerPort string
}

// helper function
func extractENV(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("Missing req env")
	}

	return val, nil
}

func Load() (Config, error) {
	// godotenv.Load() ---->> it reads .env and sets them into the process env
	// os.Getenv() ---->> reads those value
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to Load dotenv")
	}

	// extract mongodb url
	mongoURL, err := extractENV("MONGODB_URL")
	if err != nil {
		return Config{}, err
	}

	// extract mongodb database name
	mongoDB, err := extractENV("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	// extract server port
	port, err := extractENV("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURL: mongoURL,
		MongoDB: mongoDB,
		ServerPort: port,
	},nil
}
