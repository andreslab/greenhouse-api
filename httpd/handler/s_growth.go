package handler

import (
	"net/http"
	"smartcrop/platform/growth"

	"github.com/gin-gonic/gin"
)

type GrowthPostRequest struct {
	Value   string `json:"value"`
	Created string `json:"created"`
}

func GrowthPost(_growth growth.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := GrowthPostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := growth.Item{
			Value:   requestBody.Value,
			Created: "9-02-2020",
		}

		_growth.Add(item)
		c.Status(http.StatusOK)
	}
}

func GrowthGet(_growth growth.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _growth.GetAll()
		c.JSON(http.StatusOK, results)
	}

}
