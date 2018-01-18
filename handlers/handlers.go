package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateCredentials(context echo.Context) error {

	return context.JSON(http.StatusOK, "success")
}

func RetrieveCredentials(context echo.Context) error {

	return context.JSON(http.StatusOK, "success")
}

func UpdateCredentials(context echo.Context) error {

	return context.JSON(http.StatusOK, "success")
}

func DeleteCredentials(context echo.Context) error {

	return context.JSON(http.StatusOK, "success")
}
