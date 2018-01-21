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

	result, err := dynamoDB.AddEntry(&entry)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, result)
}

func RetrieveCredentials(context echo.Context) error {

	result, err := dynamoDB.GetEntry(context.Param("hostname"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, result)
}

func DeleteCredentials(context echo.Context) error {

	err := dynamoDB.DeleteEntry(context.Param("hostname"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, "success")
}
