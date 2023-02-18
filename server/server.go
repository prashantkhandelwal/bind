package server

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/bind/config"
	"github.com/prashantkhandelwal/bind/server/handlers"
)

func Run() {

	if _, err := os.Stat("data"); err != nil {
		log.Println("\"data\" directory not found!....Creating")
		if err := os.MkdirAll("data\\images", os.ModePerm); err != nil {
			log.Fatalf("ERROR: Cannot create \"data\" directory - %v", err.Error())
			panic(err)
		}
	}

	err := config.InitDB()
	if err != nil {
		log.Fatalf("ERROR: Unable to configure database - %v", err.Error())
		panic(err)
	}

	c, err := config.InitConfig()
	if err != nil {
		log.Fatalf("ERROR: Cannot load configuration - %v", err.Error())
		panic(err)
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
	router.POST("/save", handlers.SaveBookmark())

	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server! - %v", err)
	}

	log.Println("Server running!")
}
