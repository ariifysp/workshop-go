package server

import (
	_itemManagingController "github/ariifysp/workshop-go/pkg/itemManaging/controller"
	_itemManagingRepository "github/ariifysp/workshop-go/pkg/itemManaging/repository"
	_itemManagingService "github/ariifysp/workshop-go/pkg/itemManaging/service"
	_itemShopRepository "github/ariifysp/workshop-go/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemShopRepository)
	itemManagingController := _itemManagingController.NewItemManagingContollerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
	router.DELETE("/:itemID", itemManagingController.Deleting)
}
