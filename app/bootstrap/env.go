package bootstrap

import (
	"github.com/spf13/viper"
)

type Env struct {
	ENVIRONMENT       string `mapstructure:"ENVIRONMENT"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout    int    `mapstructure:"CONTEXT_TIMEOUT"`
	WikimediaUserName string `mapstructure:"WIKIMEDIA_USERNAME"`
	WikimediaPassword string `mapstructure:"WIKIMEDIA_PASSWORD"`
}

func NewEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)

	return
}
