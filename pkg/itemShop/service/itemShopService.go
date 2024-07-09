package service

import (
	_itemShopModel "github/ariifysp/workshop-go/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
