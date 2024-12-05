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
		"iss":       "banana",
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Minute * 20).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(sign))
}

func VerifyJWTToken(sign string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		// get the token from the header
		jwtValue := ctx.GetHeader("Authorization")
		if jwtValue == "" || len(jwtValue) < 7 || jwtValue[:7] != "Bearer " {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// extract the token
		tokenString := jwtValue[7:] // strings.ReplaceAll(jwtValue, "Bearer ", "")
		// parse the token
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
			elapsed := time.Since(start)
			log.Printf("took %s", elapsed)
			if elapsed < 1*time.Second {
				fmt.Println("sleeping")
				time.Sleep(1*time.Second - elapsed)
			}
			fmt.Println("responding")
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
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(p.secret), nil
}
