package handler

import (
	"net/http"
	"smartcrop/platform/plague"

	"github.com/gin-gonic/gin"
)

type PlaguePostRequest struct {
	Value   string `json:"value"`
	Created string `json:"created"`
}

//func plaguePost(_plague *plague.Repo) gin.HandlerFunc {
func PlaguePost(_plague plague.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PlaguePostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := plague.Item{
			Value:   requestBody.Value,
			Created: "9-02-2020",
		}

		_plague.Add(item)
		c.Status(http.StatusOK)
	}
}

func PlagueGet(_plague plague.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _plague.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
