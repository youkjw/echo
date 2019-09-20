package web

import (
	"echo/model"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func HomeHandler(c *echo.Context) error{
	var Banner model.Banner
	banner := Banner.GetBannerList()
	fmt.Println(banner)

	return (*c).String(http.StatusOK, "hello home!")
}