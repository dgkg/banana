package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func main() {
	r := gin.Default()
	r.POST("/register", HandleRegister)
	r.POST("/login", HandlerTest, HandlerLogin)
	r.GET("/users/:uuid", HandlerGetUserByID)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func HandlerTest(ctx *gin.Context) {
	log.Println("test")
}

type MyEncoder struct {
	zapcore.Encoder
}

func (m *MyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	filtered := make([]zapcore.Field, 0, len(fields))
	log.Println("fields:", fields)
	for _, field := range fields {
		log.Println("field key:", field.Key)
		if field.Key == "pass" || field.Key == "password" {
			continue
		}
		filtered = append(filtered, field)
	}
	return m.Encoder.EncodeEntry(entry, filtered)
}
