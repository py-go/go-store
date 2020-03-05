package main

import (
	"go-store/cmd/handlers/up"
	"go-store/cmd/router"

	"log"
	"net/http"
	"os"
)

func main() {
	// Start the Healthz service
	go healthz()

	router := router.SetupRouter()
	up.AddUpV1(router)

	log.Printf("Starting go-store API")

	port := GetServicePort("PORT", ":8080")
	log.Printf("Serving at %s", port)
	router.Run(port)

}

func GetServicePort(key, defaultPort string) string {
	if probePort := os.Getenv(key); probePort != "" {
		return ":" + probePort
	}
	return defaultPort

}
func healthz() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(GetServicePort("PROBE_PORT", ":9001"), nil)
}
