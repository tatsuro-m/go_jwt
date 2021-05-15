package auth

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
	"os"
	"time"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// ヘッダー部分
	token := jwt.New(jwt.SigningMethodHS256)

	// claims 部分
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "234543234y"
	claims["name"] = "taro"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 署名部分
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
