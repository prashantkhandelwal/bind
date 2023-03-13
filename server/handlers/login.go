package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func CreateToken(userId uint64) (string, error) {
	var err error

	access_secret := "hgfksjadhgkshfklghjlskjd"
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(access_secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login() gin.HandlerFunc {
	fn := func(g *gin.Context) {

		var u User
		if err := g.ShouldBind(&u); err != nil {
			g.JSON(http.StatusUnprocessableEntity, "Invalid JSON provide")
			return
		}

		if user.Username != u.Username || user.Password != u.Password {
			g.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}

		token, err := CreateToken(user.ID)
		if err != nil {
			g.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		cookie, err := g.Cookie("bind_cookie_status")
		if err != nil {
			cookie = "login"
			g.SetCookie("bind_cookie_status", token, 3600, "/", "localhost", false, true)
		} else {
			cookie = "login"
			g.SetCookie("bind_cookie_status", token, 3600, "/", "localhost", false, true)
		}

		g.JSON(http.StatusOK, cookie)
	}

	return gin.HandlerFunc(fn)
}
