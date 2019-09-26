package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 //8MB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		fg, err := c.FormFile("fg")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		bg, err := c.FormFile("bg")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := c.SaveUploadedFile(fg, "./fg.jpg"); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := c.SaveUploadedFile(bg, "./bg.jpg"); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		WhiteToBG("fg.jpg", "bg.jpg")

		c.Writer.Header().Set("content-type", "image/jpeg")
		http.ServeFile(c.Writer, c.Request, "./output.jpg")
	})

	router.Run(":8080")
}
