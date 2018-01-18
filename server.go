package main

import (
	"log"
	"net/http"

	db "github.com/byuoitav/pi-credentials-microservice/dynamoDB"
	km "github.com/byuoitav/pi-credentials-microservice/kms"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/pi-credentials-microservice/handlers"
	"github.com/fatih/color"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const PORT = ":5002"

func main() {

	log.Printf("%s", color.HiGreenString("Starting pi credentials microservice"))

	//start AWS clients
	awsSession := session.Must(session.NewSession())
	db.Init(session)
	km.Init(session)

	//start web server
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	//health endpoint
	secure.GET("/health", Health)

	//credential CRUD
	secure.POST("/devices/:hostname", handlers.CreateCredentials)
	secure.GET("/devices/:hostname", handlers.RetrieveCredentials)
	secure.PUT("/devices/:hostname", handlers.UpdateCredentials)
	secure.DELETE("/devices/:hostname", handlers.DeleteCredentials)

	server := http.Server{
		Addr:           PORT,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}

func Health(context echo.Context) error {

	return context.JSON(http.StatusOK, "hello from cMonster")

}
