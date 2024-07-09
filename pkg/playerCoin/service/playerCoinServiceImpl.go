package service

import (
	_playerCoinRepository "github/ariifysp/workshop-go/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}
