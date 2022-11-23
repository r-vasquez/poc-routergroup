package main

import (
	"log"
	"net/http"

	"github.com/r-vasquez/sidecar/pkg/httpapi"
	"github.com/r-vasquez/sidecar/pkg/router"
)

func main() {
	r := router.New()
	v1 := r.NewGroup("/v1/diagnostics")

	// Example of adding a top Level commands
	v1.GET("/directory", httpapi.GetDirectory)

	// Example of adding multiple groups
	controller := v1.NewGroup("/controller")
	{
		controller.GET("/cluster", httpapi.GetControllerCluster)
		controller.GET("/topic", httpapi.GetControllerTopic)
		controller.GET("/features", httpapi.GetControllerFeatures)
	}

	host := v1.NewGroup("/host")
	{
		host.GET("/cpu", httpapi.GetHostCPU)
		host.GET("/buses", httpapi.GetHostBuses)
		host.GET("/memory", httpapi.GetHostMemory)
	}

	log.Fatal(http.ListenAndServe(":8080", r.Router()))
}
