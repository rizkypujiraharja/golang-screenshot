package controller

import (
	"golang-screenshot/helper"
	"golang-screenshot/model/web"
	"golang-screenshot/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type ScreenshotControllerImpl struct {
	ScreenshotService service.ScreenshotService
}

func NewScreenshotController(screenshotService service.ScreenshotService) ScreenshotController {
	return &ScreenshotControllerImpl{
		ScreenshotService: screenshotService,
	}
}

func (controller *ScreenshotControllerImpl) Take(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	screenshotCreateRequest := web.ScreenshotTakeRequest{}
	helper.ReadFromRequestBody(request, &screenshotCreateRequest)

	createResponse := controller.ScreenshotService.Take(request.Context(), screenshotCreateRequest)
	webRespponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   createResponse,
	}

	writer.WriteHeader(http.StatusCreated)

	helper.WriteToResponseBody(writer, webRespponse)
}