package service

import (
	"context"
	"golang-screenshot/model/web"
)

type ScreenshotService interface {
	Take(ctx context.Context, request web.ScreenshotTakeRequest) web.ScreenshotResponse
}

