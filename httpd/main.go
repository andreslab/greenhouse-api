package main

import (
	"smartcrop/httpd/handler"
	"smartcrop/platform/growth"
	"smartcrop/platform/light"
	"smartcrop/platform/ph"
	"smartcrop/platform/plague"
	"smartcrop/platform/temperature"
	"smartcrop/platform/volume"
	"smartcrop/platform/wet"

	"github.com/gin-gonic/gin"
)

func main() {
	temperature := temperature.New()
	wet := wet.New()
	ph := ph.New()
	volume := volume.New()
	plague := plague.New()
	growth := growth.New()
	light := light.New()

	r := gin.Default()

	r.POST("/api/temperature", handler.TemperaturePost(temperature))
	r.POST("/api/wet", handler.WetPost(wet))
	r.POST("/api/ph", handler.PhPost(ph))
	r.POST("/api/volume", handler.VolumePost(volume))
	r.POST("/api/camera/plague", handler.PlaguePost(plague))
	r.POST("/api/canera/growth", handler.GrowthPost(growth))
	r.POST("/api/light", handler.LightPost(light))

	r.GET("/api/temperature", handler.TemperatureGet(temperature))
	r.GET("/api/wet", handler.WetGet(wet))
	r.GET("/api/ph", handler.PhGet(ph))
	r.GET("/api/volume", handler.VolumeGet(volume))
	r.GET("/api/camera/plague", handler.PlagueGet(plague))
	r.GET("/api/canera/growth", handler.GrowthGet(growth))
	r.GET("/api/light", handler.LightGet(light))

	//r.POST("/api/record", handler.NewsFeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
