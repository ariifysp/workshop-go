package controller

import "github.com/labstack/echo/v4"

type ItemShopContoller interface {
	Listing(pctx echo.Context) error
}
