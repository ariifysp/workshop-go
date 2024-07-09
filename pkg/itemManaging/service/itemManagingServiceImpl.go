package service

import (
	_itemManagingRepository "github/ariifysp/workshop-go/pkg/itemManaging/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}
