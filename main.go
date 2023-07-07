package main

import (
	"fmt"
	"net/http"

	"golang-screenshot/app"
	"golang-screenshot/controller"
	"golang-screenshot/helper"
	"golang-screenshot/middleware"
	"golang-screenshot/service"

	"github.com/go-playground/validator/v10"
)

func main()  {
	validate := validator.New()

	screenshotService := service.NewScreenshotService(validate)
	screenshotController := controller.NewScreenshotController(screenshotService)

	router := app.NewRouter(screenshotController)

	server := http.Server{
		Addr: ":8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server running on port 8080")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}