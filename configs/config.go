package configs

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App      `mapstructure:"app"`
	Postgres `mapstructure:"postgres"`
}

type App struct {
	Debug      bool   `mapstructure:"debug"`
	Env        string `mapstructure:"env"`
	Port       string `mapstructure:"port"`
	Encryptkey string `mapstructure:"encryptkey"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"database"`
}

var config Config

func InitViper(path, env string) {
	switch env {
	case "local":
		viper.SetConfigName("local-config")
	case "develop":
		viper.SetConfigName("develop-config")
	default:
		viper.SetConfigName("develop-config")
	}
	log.Println("running on environment: ", env)
	viper.AddConfigPath(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file has changed: ", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetViper() *Config {
	return &config
}
