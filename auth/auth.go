package auth

import (
    "os"
    "time"

	"github.com/gin-gonic/gin"
    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
)

func GetTokenHandler(c *gin.Context) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	// 実際にはDBから取得することを想定
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "kosuketaniguchi"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	c.JSON(200, tokenString)
}

func CheckToken() *jwtmiddleware.JWTMiddleware {
	checkToken := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNINGKEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return checkToken
}