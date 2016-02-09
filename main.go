package main

import (
	"flag"
	"github.com/cytoscape-ci/anna/reg"
	"github.com/gin-gonic/gin"
	"os"
)

func getGinRouter() *gin.Engine {
	router := gin.Default()
	setRoutes(router)
	return router
}

func setRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

func main() {
	regFile := flag.String("reg", "./registration.json", "Path to the registration file.")
	reg.Send(*regFile, os.Args[1])
	router := getGinRouter()
	router.Run(":8888")
}
