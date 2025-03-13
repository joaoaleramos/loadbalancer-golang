package main

import (
	"loadbalancer/internal"
	"loadbalancer/internal/logger"
	"loadbalancer/internal/metric"
	"loadbalancer/internal/util"

	"net/http"
	"os"
)

func main() {
	configPath := "loadfig.yaml"

	lbConfig, err := util.ParseYAML(configPath)
	if err != nil {
		logger.Logger.Fatalf("Failed to parse config: %v", err)
	}

	lb := internal.NewLoadBalancer(lbConfig.Backends)

	// Expose your metrics at /metrics
	http.Handle("/metrics", metric.MetricsHandler())
	// Handle proxy requests
	http.HandleFunc("/", lb.ServeProxy)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Logger.Info("Load Balancer running on port: ", port)
	logger.Logger.Fatal(http.ListenAndServe(":"+port, nil))
}
