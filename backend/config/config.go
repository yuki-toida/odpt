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
	DB          struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		Pool int    `toml:"pool"`
	}
	APIServer struct {
		Port string `toml:"port"`
	}
}

// Init func
func Init(configPath, envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	env := os.Getenv("ENV")
	_, err := toml.DecodeFile(configPath+"config."+env+".toml", &Config)
	if err != nil {
		panic(err)
	}

	Config.Env = env
	Config.ConsumerKey = os.Getenv("CONSUMER_KEY")
	Config.OdptURL = "https://api-tokyochallenge.odpt.org/api/v4"

	fmt.Printf("[Config] %+v\n", Config)
}
