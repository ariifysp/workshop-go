package repository

import (
	"github/ariifysp/workshop-go/entities"
	_itemManagingModel "github/ariifysp/workshop-go/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
	Deleting(itemID uint64) error
}
