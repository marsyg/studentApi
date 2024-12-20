package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP_SERVER struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTP_SERVER `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {

		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("File does not exist:%s ", configPath)
	}
	var cfg Config
	//Read Config parses from the config file
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("config path not set :%s", err.Error())
	}
	return &cfg
}
