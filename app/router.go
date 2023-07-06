package app

import (
	"golang-screenshot/controller"
	"golang-screenshot/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(screenshotController controller.ScreenshotController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/screenshot", screenshotController.Take)

	router.PanicHandler = exception.ErrorHandler

	return router
}