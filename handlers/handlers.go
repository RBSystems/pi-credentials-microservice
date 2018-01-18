package handlers

import (
	"net/http"

	"github.com/byuoitav/pi-credentials-microservice/dynamoDB"
	"github.com/byuoitav/pi-credentials-microservice/structs"
	"github.com/labstack/echo"
)

func CreateCredentials(context echo.Context) error {

	var entry structs.Entry
	err := context.Bind(&entry)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = dynamoDB.AddEntry(&entry)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, entry)
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
