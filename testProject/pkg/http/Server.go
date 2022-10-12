package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	service "testProject/internal/services"
	handler "testProject/internal/transport/rest"
)

type ServerConfig struct {
	Port       string `envconfig:"PORT" default:":8087"`
	JwkUrl     string `envconfig:"JWKS_URL" required:"true"`
	AuthJwkUrl string `envconfig:"AUTH_JWKS_URL" required:"true"`
	SkipOAuth  bool   `envconfig:"SKIP_OAUTH" required:"false" default:"false"`
}

func StartServer(serverConfig *ServerConfig, operationService service.OperationService) error {
	engine := setupServer(serverConfig, operationService)

	return engine.Run(serverConfig.Port)
}

func setupServer(
	serverConfig *ServerConfig,
	operationService service.OperationService,
) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	ginConfig := gin.LoggerConfig{SkipPaths: []string{"/"}}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"*"}

	r.Use(cors.New(corsConfig))

	r.Static("ui", "./ui-frontend")
	r.GET("/", gin.LoggerWithConfig(ginConfig), okHandler)

	// Dummy Server for tests
	if !serverConfig.SkipOAuth {
		// ToDo
	}

	operationHandler := handler.NewOperationHandler(operationService)

	r.GET("/:computerId", operationHandler.ReadHandler)
	r.POST("/update/", operationHandler.PostHandler)
	r.POST("/add/", operationHandler.PostHandler)
	r.DELETE("/delete/:computerId", operationHandler.DeleteHandler)

	return r

}

func okHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "It's our beautiful user interface!")
}
