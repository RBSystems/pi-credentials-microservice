package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/fatih/color"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const PORT = ":5002"

func main() {

	fmt.Printf("%s", color.HiGreenString("        __  ___                 __           "))
	fmt.Printf("%s", color.HiGreenString("  _____/  |/  /___  ____  _____/ /____  _____"))
	fmt.Printf("%s", color.HiGreenString(" / ___/ /|_/ / __ \\/ __ \\/ ___/ __/ _ \\/ ___/"))
	fmt.Printf("%s", color.HiGreenString("/ /__/ /  / / /_/ / / / (__  ) /_/  __/ /    "))
	fmt.Printf("%s", color.HiGreenString("\\___/_/  /_/\\____/_/ /_/____/\\__/\\___/_/     "))

	log.Printf("%s", color.HiGreenString("Starting room designation microservice..."))

	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	secure.GET("/*", Health)

	server := http.Server{
		Addr:           PORT,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}

func Health(context echo.Context) {

	return context.JSON(http.StatusOK, "hello from cMonster")

}
