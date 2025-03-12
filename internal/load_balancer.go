package internal

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
)

type Backend struct {
	URL    *url.URL `yaml:"url"`
	Weight float64  `yaml:"weight"`
}

type LoadBalancer struct {
	backends []Backend `yaml:"backends"`
}

func NewLoadBalancer(servers map[string]float64) *LoadBalancer {
	var backends []Backend
	for server, weight := range servers {
		parsedURL, _ := url.Parse(server)
		backends = append(backends, Backend{URL: parsedURL, Weight: weight})
	}

	return &LoadBalancer{
		backends: backends,
	}
}

func (lb *LoadBalancer) getNextServer() *url.URL {
	randomVal := rand.Float64()
	currentWeight := 0.0

	for _, backend := range lb.backends {
		currentWeight += backend.Weight
		if randomVal <= currentWeight {
			return backend.URL
		}
	}

	return lb.backends[len(lb.backends)-1].URL
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextServer()
	proxyURL := fmt.Sprintf("%s%s", target.String(), r.URL.Path)

	fmt.Printf("Redirecting to: %s\n", target.String())

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
