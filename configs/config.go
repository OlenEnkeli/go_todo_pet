package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type CommonSettings struct {
	Mode string `mapstructure:"COMMON_MODE"`
}

type AuthSettings struct {
	Salt         string `mapstructure:"AUTH_SALT"`
	JWTSecretKey string `mapstucture:"AUTH_JWT_SECRET_KEY"`
}

type ServerConfig struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DB       string `mapstructure:"POSTGRES_DB"`
	SSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
}

type ConfigStruct struct {
	Server   ServerConfig   `mapstructure:",squash"`
	Common   CommonSettings `mapstructure:",squash"`
	Auth     AuthSettings   `mapstructure:",squash"`
	Postgres PostgresConfig `mapstructure:",squash"`
}

func (config *ConfigStruct) GetPostgresDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.DB,
	)
}

var Config ConfigStruct

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Unable to load config: %s", err)
	}

	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("Unable to load config: %s", err)
	}
}
