package environments

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Environment      = ""
	Port             = 0
	ConnectionString = ""
	BootstrapServer  = ""
	GroupId          = ""
	Topic            = ""
	SlackBaseUrl     = ""
)

func NewConfig() {
	var err error

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	Environment = viper.GetString("environment")
	Port = viper.GetInt("api.port")
	ConnectionString = viper.GetString("api.connectionString")
	BootstrapServer = viper.GetString("consumer.bootstrapServer")
	GroupId = viper.GetString("consumer.groupId")
	Topic = viper.GetString("consumer.topic")
	SlackBaseUrl = viper.GetString("consumer.slackBaseUrl")
}
