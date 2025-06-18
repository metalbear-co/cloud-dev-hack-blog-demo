package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

var (
	internalServiceHost = getEnv("INTERNAL_SERVICE_HOST", "internal-api.default.svc.cluster.local")
	internalServiceURL  = fmt.Sprintf("http://%s:8080/process", internalServiceHost)
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("Pinging internal service at: %s", internalServiceURL)
			resp, err := http.Get(internalServiceURL)
			if err != nil {
				log.Printf("Error: %v", err)
				continue
			}
			defer resp.Body.Close()
			log.Printf("Response from internal service: %s", resp.Status)
		}
	}
}
