package main

import (
	"loadbalancer/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	servers := map[string]float64{
		"http://localhost:5001": 0.7,
		"http://localhost:5002": 0.3,
	}

	lb := internal.NewLoadBalancer(servers)

	http.HandleFunc("/", lb.ServeProxy)

	port := os.Getenv("PORT")
	log.Println("Load Balancer running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
