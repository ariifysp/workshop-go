package repository

type adminRepositoryImpl struct{}

func NewAdminRepositoryImpl() AdminRepository {
	return &adminRepositoryImpl{}
}
