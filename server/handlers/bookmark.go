package handlers

import (
	"log"
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
		id, err := db.Save(&bookmark)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if id > 0 {
			g.JSON(http.StatusOK, gin.H{"result": "OK"})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong in saving the bookmark."})
		}

		if id > 0 {
			log.Println(bookmark)
			go db.GetSnap(id, bookmark.Url)
		}
	}
	return gin.HandlerFunc(fn)
}
