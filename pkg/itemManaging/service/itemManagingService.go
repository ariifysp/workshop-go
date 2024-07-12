package service

import (
	_itemCreatingModel "github/ariifysp/workshop-go/pkg/itemManaging/model"
	_itemShopModel "github/ariifysp/workshop-go/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemCreatingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemCreatingModel.ItemEditingReq) (*_itemShopModel.Item, error)
	Deleting(itemID uint64) error
}
