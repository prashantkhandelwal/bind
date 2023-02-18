package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/bind/webext"
)

func WebExtract() gin.HandlerFunc {
	fn := func(g *gin.Context) {

		url := g.Query("url")

		if len(strings.TrimSpace(url)) > 0 {
			w, err := webext.Extract(url)
			if err != nil {
				log.Fatalf("ERROR: %s", err.Error())
			}
			g.JSON(http.StatusOK, gin.H{"result": &w})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"result": "Error: Url cannot be blank."})
		}
	}
	return gin.HandlerFunc(fn)
}
