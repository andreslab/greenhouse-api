package handler

import (
	"net/http"
	"smartcrop/platform/volume"

	"github.com/gin-gonic/gin"
)

type VolumePostRequest struct {
	Value   string `json:"value"`
	Tank    int    `json:"tank"`
	Created string `json:"created"`
}

//func volumePost(_volume *volume.Repo) gin.HandlerFunc {
func VolumePost(_volume volume.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := VolumePostRequest{}
		c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		item := volume.Item{
			Value:   requestBody.Value,
			Tank:    requestBody.Tank,
			Created: "9-02-2020",
		}

		_volume.Add(item)
		c.Status(http.StatusOK)
	}
}

func VolumeGet(_volume volume.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := _volume.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
