package util

import (
	"loadbalancer/internal"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseYAML(path string) (internal.LoadBalancer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return internal.LoadBalancer{}, err
	}
	var config internal.LoadBalancer
	if err := yaml.Unmarshal(data, &config); err != nil {
		return internal.LoadBalancer{}, err
	}
	return config, nil
}
