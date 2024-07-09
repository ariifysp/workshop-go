package repository

type itemManagingRepositoryImpl struct{}

func NewItemManagingRepositoryImpl() ItemManagingRepository {
	return &itemManagingRepositoryImpl{}
}
