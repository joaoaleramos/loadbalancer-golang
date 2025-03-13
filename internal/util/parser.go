package util

import (
	"loadbalancer/internal"
	"loadbalancer/internal/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type LoadBalancerConfiguration struct {
	Services []internal.Backend `yaml:"services"`
}

func ParseYAML(path string) (internal.LoadBalancer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return internal.LoadBalancer{}, err
	}
	var config LoadBalancerConfiguration
	if err := yaml.Unmarshal(data, &config); err != nil {
		return internal.LoadBalancer{}, err
	}
	logger.Logger.Info("Parsed config:", config)
	return internal.LoadBalancer{Backends: config.Services}, nil
}
