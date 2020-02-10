package handler

import (
	"net/http"
	"smartcrop/platform/temperature"

	"github.com/gin-gonic/gin"
)

type TemperaturePostRequest struct {
	Value   string `json:"value"`
	Zone    string `jsone:"zone"`
	Created string `json:"created"`
}

//func temperaturePost(_temperature *temperature.Repo) gin.HandlerFunc {
func TemperaturePost(_temperature temperature.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := TemperaturePostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := temperature.Item{
			Value:   requestBody.Value,
			Zone:    requestBody.Zone,
			Created: "9-02-2020",
		}

		_temperature.Add(item)
		c.Status(http.StatusOK)
	}
}

func TemperatureGet(_temperature temperature.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _temperature.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
