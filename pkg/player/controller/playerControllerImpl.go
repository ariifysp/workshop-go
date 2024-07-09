package controller

import (
	_playerService "github/ariifysp/workshop-go/pkg/player/service"
)

type playerControllerImpl struct {
	playerService _playerService.PlayerService
}

func NewPlayerControllerImpl(
	playerService _playerService.PlayerService,
) PlayerController {
	return &playerControllerImpl{playerService}
}
