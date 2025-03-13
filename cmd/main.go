package main

import (
	"loadbalancer/internal"
	"loadbalancer/internal/logger"
	"loadbalancer/internal/util"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func main() {
	configPath := "loadfig.yaml"
	lbConfig, err := util.ParseYAML(configPath)
	if err != nil {
		logger.Logger.Fatal("Failed to parse config: %v", zap.Error(err))
	}

	lb := internal.NewLoadBalancer(lbConfig.Backends)

	http.HandleFunc("/", lb.ServeProxy)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Logger.Info("Load Balancer running on port:", zap.String("port", port))
	logger.Logger.Fatal(http.ListenAndServe(":"+port, nil).Error())
}
