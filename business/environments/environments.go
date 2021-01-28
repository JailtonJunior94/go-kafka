package environments

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Environment      = ""
	ConnectionString = ""
	BootstrapServer  = ""
	GroupId          = ""
	Topic            = ""
	Port             = 0
	SlackBaseUrl     = ""
)

func NewConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ConnectionString = os.Getenv("CONNECTION_STRING")
	BootstrapServer = os.Getenv("BOOTSTRAP_SERVER")
	GroupId = os.Getenv("GROUP_ID")
	Topic = os.Getenv("TOPIC")
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}
	SlackBaseUrl = os.Getenv("SLACK_BASEURL")
}
