package handler

import (
	"net/http"
	"smartcrop/platform/wet"

	"github.com/gin-gonic/gin"
)

type WetPostRequest struct {
	Value   string `json:"value"`
	Zone    string `jsone:"zone"`
	Created string `json:"created"`
}

//func wetPost(_wet *wet.Repo) gin.HandlerFunc {
func WetPost(_wet wet.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := WetPostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := wet.Item{
			Value:   requestBody.Value,
			Zone:    requestBody.Zone,
			Created: "9-02-2020",
		}

		_wet.Add(item)
		c.Status(http.StatusOK)
	}
}

func WetGet(_wet wet.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _wet.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
