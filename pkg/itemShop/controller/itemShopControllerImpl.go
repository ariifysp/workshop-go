package controller

import (
	"github/ariifysp/workshop-go/pkg/custom"
	_itemShopModel "github/ariifysp/workshop-go/pkg/itemShop/model"
	_itemShopService "github/ariifysp/workshop-go/pkg/itemShop/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopService.ItemShopService,
) ItemShopContoller {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
