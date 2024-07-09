package controller

import (
	_inventoryService "github/ariifysp/workshop-go/pkg/inventory/service"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
}

func NewInventoryControllerImpl(
	inventoryService _inventoryService.InventoryService,
) InventoryController {
	return &inventoryControllerImpl{inventoryService}
}
