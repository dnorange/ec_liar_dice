package main

import (
	"fmt"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/dnorange/ec_liar_dice/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()

	if err != nil {
		fmt.Println("cf env error")
	}

	server := service.NewServer(appEnv)
	server.Run(":" + port)
}
