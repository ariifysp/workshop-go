package service

import (
	_inventoryRepository "github/ariifysp/workshop-go/pkg/inventory/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
}

func NewInventoryServiceImpl(
	inventoryRepository _inventoryRepository.InventoryRepository,
) InventoryService {
	return &inventoryServiceImpl{inventoryRepository}
}
