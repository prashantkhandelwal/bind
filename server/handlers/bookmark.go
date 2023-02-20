package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prashantkhandelwal/bind/db"
)

func SaveBookmark() gin.HandlerFunc {
	fn := func(g *gin.Context) {

		var bookmark db.Bookmark
		if err := g.BindJSON(&bookmark); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		save, err := db.Save(&bookmark)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if save == true {
			g.JSON(http.StatusOK, gin.H{"result": "OK"})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong in saving the bookmark."})
		}
	}
	return gin.HandlerFunc(fn)
}
