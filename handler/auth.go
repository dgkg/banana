package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func BasicAuth(ctx *gin.Context) {
	user, pass, ok := ctx.Request.BasicAuth()
	if !ok || user != "admin" || pass != "admin" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Next()
}

func (h *Handler) newJWTToken(uuidUser, sign string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid_user": uuidUser,
		"issuer":    "banana",
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Minute * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(sign))
}

func VerifyJWTToken(sign string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtValue := ctx.GetHeader("Authorization")
		if jwtValue == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if len(jwtValue) < 7 || jwtValue[:7] != "Bearer " {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//
		tokenString := jwtValue[7:]
		log.Println("tokenString:", tokenString)
		checkJWT := &parseMethod{secret: sign}
		token, err := jwt.Parse(tokenString, checkJWT.parser)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println("token is valid", claims["iat"], claims["exp"], claims["uuid_user"])
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

type parseMethod struct {
	secret string
}

func (p *parseMethod) parser(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(p.secret), nil
}
