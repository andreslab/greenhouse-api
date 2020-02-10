package handler

import (
	"net/http"
	"smartcrop/platform/ph"

	"github.com/gin-gonic/gin"
)

type PhPostRequest struct {
	Value   string `json:"value"`
	Tank    int    `jsone:"tank"`
	Created string `json:"created"`
}

//func phPost(_ph *ph.Repo) gin.HandlerFunc {
func PhPost(_ph ph.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PhPostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := ph.Item{
			Value:   requestBody.Value,
			Tank:    requestBody.Tank,
			Created: "9-02-2020",
		}

		_ph.Add(item)
		c.Status(http.StatusOK)
	}
}

func PhGet(_ph ph.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _ph.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
