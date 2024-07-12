package repository

import (
	"github/ariifysp/workshop-go/databases"
	"github/ariifysp/workshop-go/entities"

	"github.com/labstack/echo/v4"

	_itemManagingException "github/ariifysp/workshop-go/pkg/itemManaging/exception"
	_itemManagingModel "github/ariifysp/workshop-go/pkg/itemManaging/model"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Deleting(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_deleted", true).Error; err != nil {
		r.logger.Errorf("Deleting item failed: %s", err.Error())
		return &_itemManagingException.ItemDeleting{ItemID: itemID}
	}

	return nil
}
