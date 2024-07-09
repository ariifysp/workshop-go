package controller

import (
	_itemManagingService "github/ariifysp/workshop-go/pkg/itemManaging/service"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingServiceImpl(
	itemManagingService _itemManagingService.ItemManagingService,
) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}
