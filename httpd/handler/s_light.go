package handler

import (
	"net/http"
	"smartcrop/platform/light"

	"github.com/gin-gonic/gin"
)

type LightPostRequest struct {
	Value   string `json:"value"`
	Created string `json:"created"`
}

//func lightPost(_light *light.Repo) gin.HandlerFunc {
func LightPost(_light light.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := LightPostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := light.Item{
			Value:   requestBody.Value,
			Created: "9-02-2020",
		}

		_light.Add(item)
		c.Status(http.StatusOK)
	}
}

func LightGet(_light light.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _light.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
