package inventoryHandler

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"

type (
	inventoryGrpcHandler struct {
		inventoryService inventoryService.InventoryService
	}
)

func NewInventoryGrpcHandler(inventoryService inventoryService.InventoryService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{inventoryService}
}
