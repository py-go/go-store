package main

import (
	"go-store/cmd/handlers/auth"
	"go-store/cmd/handlers/product"
	"go-store/cmd/handlers/up"
	"go-store/cmd/middleware"
	"go-store/cmd/router"
	"go-store/cmd/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Start the Healthz service
	go healthz()

	router := router.SetupRouter()
	router.Use(middleware.SetUserStatus())
	router.NoRoute(handle404)
	router.LoadHTMLGlob("templates/*")
	product.AddProduct(router)
	up.AddUpV1(router)
	auth.AddAuth(router)

	log.Printf("Starting go-store API")

	port := GetServicePort("PORT", ":8080")
	log.Printf("Serving at %s", port)
	router.Run(port)

}

func handle404(c *gin.Context) {
	utils.RenderError(c, http.StatusNotFound, "error.html", gin.H{
		"message": "Page not found",
		"title":   "Error",
		"code":    "PAGE_NOT_FOUND",
	})

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
