package repository

type playerRepositoryImpl struct{}

func NewPlayerRepositoryImpl() PlayerRepository {
	return &playerRepositoryImpl{}
}
