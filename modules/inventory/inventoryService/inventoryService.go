package inventoryService

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryRepository"

type (
	InventoryService interface{}

	inventoryService struct {
		inventoryRepo inventoryRepository.InventoryRepository
	}
)

func NewInventoryService(inventoryRepo inventoryRepository.InventoryRepository) InventoryService {
	return &inventoryService{inventoryRepo}
}
