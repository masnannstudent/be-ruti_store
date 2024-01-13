package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"

	"os"
	"strconv"
)

type Config struct {
	ServerPort  int
	DatabaseUrl string
	Secret      string
	ServerKey   string
	ClientKey   string
	CCName      string
	CCAPIKey    string
	CCAPISecret string
	CCFolder    string
}

func InitConfig() *Config {
	return loadConfig()

}

func loadConfig() *Config {

	var res = new(Config)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	if value, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DATABASE_URL"); found {
		res.DatabaseUrl = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val, found := os.LookupEnv("SERVERKEY"); found {
		res.ServerKey = val
	}

	if val, found := os.LookupEnv("CLIENTKEY"); found {
		res.ClientKey = val
	}

	if value, found := os.LookupEnv("CCNAME"); found {
		res.CCName = value
	}

	if value, found := os.LookupEnv("CCAPIKEY"); found {
		res.CCAPIKey = value
	}
	if value, found := os.LookupEnv("CCAPISECRET"); found {
		res.CCAPISecret = value
	}
	if value, found := os.LookupEnv("CCFOLDER"); found {
		res.CCFolder = value
	}

	return res
}
