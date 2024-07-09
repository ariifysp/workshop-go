package service

import (
	_playerRepository "github/ariifysp/workshop-go/pkg/player/repository"
)

type playerServiceImpl struct {
	playerRepository _playerRepository.PlayerRepository
}

func NewPlayerServiceImpl(
	playerRepository _playerRepository.PlayerRepository,
) PlayerService {
	return &playerServiceImpl{playerRepository}
}
