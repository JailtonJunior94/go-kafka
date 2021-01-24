package environments

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Environment      = ""
	ConnectionString = ""
	BootstrapServer  = ""
	GroupId          = ""
	Topic            = ""
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
}
