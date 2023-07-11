package service

import (
	"context"
	"encoding/base64"
	"golang-screenshot/helper"
	"golang-screenshot/model/web"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/go-playground/validator/v10"
)

type ScreenshotServiceImpl struct {
	Validate *validator.Validate
}

func NewScreenshotService (validate *validator.Validate) ScreenshotService {
	return &ScreenshotServiceImpl{
		Validate: validate,
	}
}

func (service ScreenshotServiceImpl) Take(ctx context.Context, request web.ScreenshotTakeRequest) web.ScreenshotResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	ctx_chrome, _ := chromedp.NewContext(ctx)
	// defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx_chrome,
		chromedp.EmulateViewport(int64(request.Width), int64(request.Height)),
		chromedp.Navigate(request.Url),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		log.Fatal(err)
	}

	// pngFile, err := os.Create("./shot.png")
    // if err != nil {
    //     panic(err)
    // }
    // defer pngFile.Close()

    // pngFile.Write(buf)

	sEnc := base64.StdEncoding.EncodeToString([]byte(buf))

	screenshot := web.ScreenshotResponse{
		Image: sEnc,
	}

	return screenshot
}