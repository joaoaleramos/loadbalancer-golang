package internal

import (
	"fmt"
	"loadbalancer/internal/logger"
	"loadbalancer/internal/metric"
	"math/rand/v2"
	"net/http"
	"net/url"
)

type Backend struct {
	Name   string  `yaml:"name"`
	URL    string  `yaml:"url"`
	Weight float64 `yaml:"weight"`
}

type LoadBalancer struct {
	Backends []Backend `yaml:"backends"`
}

func NewLoadBalancer(services []Backend) *LoadBalancer {
	return &LoadBalancer{Backends: services}
}

// getNextServer returns a parsed URL based on weight.
func (lb *LoadBalancer) getNextServer() *url.URL {
	if len(lb.Backends) == 0 {
		return nil
	}
	randomVal := rand.Float64()
	currentWeight := 0.0
	for _, backend := range lb.Backends {
		currentWeight += backend.Weight
		if randomVal <= currentWeight {
			parsedURL, err := url.Parse(backend.URL)
			if err != nil {
				return nil
			}
			metric.RequestsHandled.WithLabelValues(backend.Name).Inc()
			return parsedURL
		}
	}
	// Fallback to last backend
	last := lb.Backends[len(lb.Backends)-1]
	metric.RequestsHandled.WithLabelValues(last.Name).Inc()
	parsedURL, _ := url.Parse(last.URL)
	return parsedURL
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextServer()
	if target == nil {
		http.Error(w, "No backends available", http.StatusServiceUnavailable)
		return
	}
	proxyURL := fmt.Sprintf("%s%s", target.String(), r.URL.Path)
	logger.Logger.Info("Redirecting to:", target)

	resp, err := http.Get(proxyURL)
	if err != nil {
		http.Error(w, "Error while redirecting request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	resp.Write(w)
}
