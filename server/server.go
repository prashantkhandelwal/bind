package server

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/bind/config"
	"github.com/prashantkhandelwal/bind/server/handlers"
)

func Run() {

	err := config.InitDB()
	if err != nil {
		log.Fatalf("ERROR: Unable to configure database - %v", err)
	}

	c, err := config.InitConfig()
	if err != nil {
		log.Fatalf("ERROR: Cannot load configuration - %v", err)
	}

	port := c.Server.PORT

	if c.Environment != "" {
		if strings.ToLower(c.Environment) == "release" {
			gin.SetMode(gin.ReleaseMode)
		} else {
			gin.SetMode(gin.DebugMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.GET("/ping", handlers.Ping)
	router.GET("/web/extract/", handlers.WebExtract())

	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server! - %v", err)
	}

	log.Println("Server running!")
}
