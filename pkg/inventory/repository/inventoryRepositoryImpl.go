package repository

type inventoryRepositoryImpl struct{}

func NewInventoryRepositoryImpl() InventoryRepository {
	return &inventoryRepositoryImpl{}
}
