package web

import (
	"echo/model"
	"github.com/labstack/echo"
	"net/http"
)

func HomeHandler(c *echo.Context) error{
	var Banner model.Banner
	banner := Banner.GetBannerList()

	return (*c).String(http.StatusOK, "hello " + banner.Name)
}