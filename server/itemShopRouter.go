package server

import (
	_itemShopController "github/ariifysp/workshop-go/pkg/itemShop/controller"
	_itemShopRepository "github/ariifysp/workshop-go/pkg/itemShop/repository"
	_itemShopService "github/ariifysp/workshop-go/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
}
