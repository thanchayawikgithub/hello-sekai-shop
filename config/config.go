package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App      `mapstructure:"app" validate:"required"`
		DB       DB       `mapstructure:"database" validate:"required"`
		Jwt      Jwt      `mapstructure:"jwt" validate:"required"`
		Kafka    Kafka    `mapstructure:"kafka" validate:"required"`
		Grpc     Grpc     `mapstructure:"grpc" validate:"required"`
		Paginate Paginate `mapstructure:"paginate" validate:"required"`
	}

	App struct {
		Stage string `mapstructure:"stage" validate:"required"`
		Name  string `mapstructure:"name" validate:"required"`
		URL   string `mapstructure:"url" validate:"required"`
	}

	DB struct {
		URI string `mapstructure:"uri" validate:"required"`
	}

	Jwt struct {
		AccessSecretKey  string `mapstructure:"access_secret_key" validate:"required"`
		RefreshSecretKey string `mapstructure:"refresh_secret_key" validate:"required"`
		ApiSecretKey     string `mapstructure:"api_secret_key" validate:"required"`
		AccessDuration   int64  `mapstructure:"access_duration" validate:"required,gte=0"`
		RefreshDuration  int64  `mapstructure:"refresh_duration" validate:"required,gte=0"`
		ApiDuration      int64  `mapstructure:"api_duration"`
	}

	Kafka struct {
		URL       string `mapstructure:"url" validate:"required"`
		ApiKey    string `mapstructure:"api_key" validate:"required"`
		ApiSecret string `mapstructure:"api_secret" validate:"required"`
	}

	Grpc struct {
		AuthURL      string `mapstructure:"auth_url" validate:"required"`
		PlayerURL    string `mapstructure:"player_url" validate:"required"`
		ItemURL      string `mapstructure:"item_url" validate:"required"`
		InventoryURL string `mapstructure:"inventory_url" validate:"required"`
		PaymentURL   string `mapstructure:"payment_url" validate:"required"`
	}

	Paginate struct {
		ItemNextPageBasedURL      string `mapstructure:"item_next_page_based_url" validate:"required"`
		InventoryNextPageBasedURL string `mapstructure:"inventory_next_page_based_url" validate:"required"`
	}
)

// LoadConfig loads the configuration based on the environment and path
func LoadConfig(state string, service string) *Config {
	var config Config
	fmt.Println(state, service)
	viper.AddConfigPath(fmt.Sprintf("config/%s", state))
	viper.SetConfigName(fmt.Sprintf("config.%s", service))
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	validate := validator.New()

	if err := validate.Struct(config); err != nil {
		panic(err)
	}

	return &config
}
