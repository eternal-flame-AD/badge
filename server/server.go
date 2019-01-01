package main

import (
	"github.com/eternal-flame-AD/badge/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router.Router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "https://badge.eternalflame.info")
	})
	router.Router.Run(":8080")
}
