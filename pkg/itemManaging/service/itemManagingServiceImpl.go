package service

import (
	"github/ariifysp/workshop-go/entities"
	_itemManagingRepository "github/ariifysp/workshop-go/pkg/itemManaging/repository"
	_itemShopRepository "github/ariifysp/workshop-go/pkg/itemShop/repository"

	_itemCreatingModel "github/ariifysp/workshop-go/pkg/itemManaging/model"
	_itemShopModel "github/ariifysp/workshop-go/pkg/itemShop/model"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository, itemShopRepository}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemCreatingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemCreatingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	_, err := s.itemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepository.FindByID(itemID)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Deleting(itemID uint64) error {
	return s.itemManagingRepository.Deleting(itemID)
}
