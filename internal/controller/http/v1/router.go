// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "gctree/docs"
	"gctree/internal/usecase"
	"gctree/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       gctree API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8820
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.CTree, serviceName, checkApi string) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe & consul
	handler.GET(checkApi, func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newAppRoutes(h, t, l)
	}
}
