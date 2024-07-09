package repository

type playerCoinRepositoryImpl struct{}

func NewPlayerCoinRepositoryImpl() PlayerCoinRepository {
	return &playerCoinRepositoryImpl{}
}
