package main

import (
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/router"
)

// @title GEO API
// @version 2.0
// @description This is a sample API for address searching and geocoding using Dadata API.
// @host localhost:8080
// @termsOfService http://localhost:8080/swagger/index.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	metrics.MustRegister()
	r := router.New()
	r.Start()
}
