package bootstrap

import (
	"log"
	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET           string `mapstructure:"JWT_SECRET"`
	MONGO_URL            string `mapstructure:"MONGO_URL"`
	MONGO_DATABASE       string `mapstructure:"MONGO_DATABASE"`
	SERVER_ADDRESS       string `mapstructure:"SERVER_ADDRESS"`
	USER_COLLECTION      string `mapstructure:"USER_COLLECTION"`
	TASK_COLLECTION      string `mapstructure:"TASK_COLLECTION"`
	ALLOWED_USERS        string `mapstructure:"ALLOWED_USERS"`
	TEST_DATABASE        string `mapstructure:"TEST_DATABASE"`
	TEST_USER_COLLECTION string `mapstructure:"TEST_USER_COLLECTION"`
	TEST_TASK_COLLECTION string `mapstructure:"TEST_TASK_COLLECTION"`
}

func NewEnv() *Env {
	// Initialize viper to read from environment variables
	viper.AutomaticEnv()

	env := Env{}

	err := viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Environment can't be loaded : %v", err)
	}

	return &env
}
