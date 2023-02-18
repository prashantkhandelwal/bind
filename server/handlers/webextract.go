package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/bind/webext"
)

func WebExtract() gin.HandlerFunc {
	fn := func(g *gin.Context) {

		url := g.Query("url")
		w := make(chan webext.WebData)

		if len(strings.TrimSpace(url)) > 0 {
			go webext.ExtractMeta(url, w)
			g.JSON(http.StatusOK, gin.H{"result": <-w})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"result": "Error: Url cannot be blank."})
		}
	}
	return gin.HandlerFunc(fn)
}
