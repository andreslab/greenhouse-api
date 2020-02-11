package handler

import (
	"net/http"
	"smartcrop/platform/wet"
	"time"

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
		//requestBody := WetPostRequest{}
		//c.Bind(&requestBody) //guarda en requestBody la data del body, debe tener la misma estructura

		value := c.PostForm("value")
		zone := c.PostForm("zone")
		currentTime := time.Now()
		today := currentTime.Format(time.RFC850)

		item := wet.Item{
			Value:   value,
			Zone:    zone,
			Created: today,
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
