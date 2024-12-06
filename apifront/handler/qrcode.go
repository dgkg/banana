package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func (h *Handler) GetQRCode(ctx *gin.Context) {
	// solutionQRCode1(ctx)
	solutionQRCode2(ctx)
}

// solutionQRCode1 is a solution to create a qrcode and save it to local file
// then read the file and write to response.
func solutionQRCode1(ctx *gin.Context) {
	// create a qrcode
	qrc, err := qrcode.New("https://github.com/yeqown/go-qrcode")
	if err != nil {
		respError(ctx, "qrcode", err)
		return
	}

	// New will create file automatically.
	options := []standard.ImageOption{
		standard.WithBgColorRGBHex("#ffffff"),
		standard.WithFgColorRGBHex("#000000"),
		// more ...
	}

	//	create the filename
	fileNAme := uuid.New().String() + ".jpg"
	writer, err := standard.New(fileNAme, options...)
	if err != nil {
		respError(ctx, "qrcode", err)
		return
	}

	// write to response
	if err = qrc.Save(writer); err != nil {
		fmt.Printf("could not save image: %v", err)
		respError(ctx, "qrcode", err)
		return
	}

	data, err := os.ReadFile(fileNAme)
	if err != nil {
		respError(ctx, "qrcode", err)
		return
	}

	// delete the local file
	err = os.Remove(fileNAme)
	if err != nil {
		respError(ctx, "qrcode", err)
		return
	}

	// respond
	ctx.Writer.Header().Set("Content-Type", "image/jpeg")
	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write(data)
}

// solutionQRCode2 is a solution to create a qrcode and save it to response directly.
// the response writer is a custom writer which implements io.WriteCloser.
func solutionQRCode2(ctx *gin.Context) {
	// create a qrcode
	qrc, err := qrcode.New("https://github.com/yeqown/go-qrcode")
	if err != nil {
		respError(ctx, "qrcode", err)
		return
	}

	// New will create file automatically.
	options := []standard.ImageOption{
		standard.WithBgColorRGBHex("#FF00FF"),
		standard.WithFgColorRGBHex("#000000"),
	}

	// create a custom writer
	rwc := &RWCloser{ctx.Writer}

	//	create the filename
	writer := standard.NewWithWriter(rwc, options...)

	// write to response
	if err = qrc.Save(writer); err != nil {
		fmt.Printf("could not save image: %v", err)
		respError(ctx, "qrcode", err)
		return
	}

	// respond
	ctx.Writer.Header().Set("Content-Type", "image/jpeg")
	ctx.Writer.WriteHeader(200)
}

type RWCloser struct {
	gin.ResponseWriter
}

func (rwc *RWCloser) Close() error {
	notify := rwc.CloseNotify()
	val, ok := <-notify
	if !ok {
		log.Println("notify: allready close")
		return nil
	}
	log.Println("notify: ", val)
	return nil
}
