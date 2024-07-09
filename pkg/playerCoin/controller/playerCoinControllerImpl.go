package controller

import (
	_playerCoinService "github/ariifysp/workshop-go/pkg/playerCoin/service"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinServiceImpl(
	playerCoinService _playerCoinService.PlayerCoinService,
) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}
