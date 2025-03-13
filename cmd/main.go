package main

import (
	"loadbalancer/cmd/util"
	"loadbalancer/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	configPath := "loadfig.yaml"
	lbConfig, err := util.ParseYAML(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	lb := internal.NewLoadBalancer(lbConfig.Backends)

	http.HandleFunc("/", lb.ServeProxy)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Load Balancer running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
