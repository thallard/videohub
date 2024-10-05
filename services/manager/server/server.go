package server

import "github.com/gin-gonic/gin"

func setupRouter() {
	r := gin.Default()

	monitoringRoutes := r.Group("/mon")
	monitoringRoutes.GET("/healthcheck", healthCheck)

	apiRoutes := r.Group("/v1")
	apiRoutes.GET("/user/:id")
}
