package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

// Config struct
var Config struct {
	Env         string
	ConsumerKey string
	OdptURL     string
	Server      struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	}
	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
	}
}

// Init func
func Init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	env := "dev"
	if v := os.Getenv("ENV"); v != "" {
		env = v
	}
	_, err := toml.DecodeFile("config/"+env+".toml", &Config)
	if err != nil {
		panic(err)
	}

	Config.Env = env
	Config.ConsumerKey = os.Getenv("CONSUMER_KEY")
	Config.OdptURL = "https://api-tokyochallenge.odpt.org/api/v4"
	fmt.Printf("[Config] %+v\n", Config)
}
