package main

import (
	"golang-rest-api/app"
	"golang-rest-api/handler"
	"golang-rest-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

var (
	mhs = new(handler.MahasiswaHandler)
)

func main() {
	app.DatabaseConnect()

	r := gin.Default()
	defer r.Run(":8888")

	r.GET("/", func(c *gin.Context) {
		res := gin.H{
			"status":  http.StatusOK,
			"message": "REST-Api",
		}
		c.JSON(http.StatusOK, res)
	})

	mahasiswa := r.Group("/mahasiswa")
	{
		mahasiswa.GET("/", middleware.Auth, mhs.ShowAll)
		mahasiswa.GET("/:id", middleware.Auth, mhs.Show)
		mahasiswa.POST("/create", middleware.Auth, mhs.Create)
		mahasiswa.PUT("/update/:id", middleware.Auth, mhs.Update)
		mahasiswa.DELETE("/delete/:id", middleware.Auth, mhs.Delete)
	}
}
