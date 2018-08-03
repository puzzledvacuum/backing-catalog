package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	eeureka "github.com/puzzledvacuum/backing-catalog/eeureka"
	service "github.com/puzzledvacuum/backing-catalog/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	eeureka.RegisterAt("http://localhost:8081", "backing-catalog", port, "8443")

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	backend, _ := eeureka.GetServiceInstances("backing-fulfillment")
	fmt.Print(backend.HostName)

	server := service.NewServerFromCFEnv(appEnv)
	server.Run(":" + port)
}
