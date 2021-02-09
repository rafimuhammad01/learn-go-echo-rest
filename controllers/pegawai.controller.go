package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/rafimuhammad01/learn-go-echo-rest/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func FetchAllPegawai(c echo.Context) error {
	res, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func StorePegawai(c echo.Context) error{
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	res, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("_id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	var err error
	var objectId primitive.ObjectID


	objectId, err  = primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	res, err := models.UpdatePegawai(objectId,nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func DeletePegawai(c echo.Context) error  {
	id := c.FormValue("_id")

	var err error
	var objectId primitive.ObjectID

	objectId, err  = primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	res, err := models.DeletePegawai(objectId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK, res)


}