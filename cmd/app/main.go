package main

import (
	"go-store/cmd/controllers"
	"go-store/cmd/database"
	"go-store/cmd/demo"
	"go-store/cmd/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func init() {

	database.Init()

}
func cli() {
	if args := os.Args; len(args) > 1 {
		arg := args[1]
		if arg == "loaddata" {
			demo.CreateDemoData(database.Connection())
		}
		os.Exit(0)
	}
}
func main() {

	cli()
	go healthz()

	router := gin.Default()
	router.NoRoute(handle404)
	router.Use(cors.Default())
	router.Use(middlewares.SessionMiddleware())

	api := router.Group("/v1/")
	controllers.Init(api)

	log.Printf("Starting go-store API")
	port := GetServicePort("PORT", ":8080")

	log.Printf("Serving at %s", port)
	router.Run(port)

}

func handle404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
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
