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

	ch1 := make(chan bool)
	eeureka.Checkpulse("backing-fulfillment", ch1)
	_ = <-ch1

	fmt.Print(ch1)

	// ticker := time.NewTicker(500 * time.Millisecond)
	// count := 1

	// for t := range ticker.C {
	// 	frontend, err := eeureka.GetServiceInstances("backing-fulfillment")
	// 	if count%4 == 0 {
	// 		fmt.Println("...Waiting for message:", t)
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("Service not available, %v\n", err)
	// 		// return
	// 	}
	// 	if len(frontend) > 0 {
	// 		fmt.Println("HostName:", frontend[0].HostName)
	// 		fmt.Println("Port:", frontend[0].Port.Port)
	// 		break
	// 	}
	// 	count++
	// 	// if count%60 == 0 {
	// 	// 	break
	// 	// }
	// }

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	// backend, _ := eeureka.GetServiceInstances("backing-fulfillment")
	// fmt.Printf("%v\n", backend[0].HostName)
	// fmt.Printf("%v\n", backend[0].Port.Port)

	server := service.NewServerFromCFEnv(appEnv)
	server.Run(":" + port)
}
