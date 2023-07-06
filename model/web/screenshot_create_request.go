package web

type ScreenshotTakeRequest struct {
	Width int `validate:"required,min=1,max=2000" json:"width"`
	Height int `validate:"required,min=1,max=2000" json:"height"`
	Url string `validate:"required,url" json:"url"`
}