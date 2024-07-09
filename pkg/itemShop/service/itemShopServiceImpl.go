package service

import (
	"github/ariifysp/workshop-go/entities"
	_itemShopModel "github/ariifysp/workshop-go/pkg/itemShop/model"
	_itemShopRepository "github/ariifysp/workshop-go/pkg/itemShop/repository"
	"math"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	page := itemFilter.Page
	size := itemFilter.Size
	totalPage := math.Ceil(float64(itemCounting) / float64(size))

	result := s.toItemResultResponse(itemList, page, int64(totalPage))

	return result, nil
}

func (s *itemShopServiceImpl) toItemResultResponse(itemsEntitiesList []*entities.Item, page int64, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemsEntitiesList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
