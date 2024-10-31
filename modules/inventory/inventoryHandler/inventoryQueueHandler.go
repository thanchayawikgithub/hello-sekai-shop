package inventoryHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"
)

type (
	InventoryQueueHandler interface {
	}

	inventoryQueueHandler struct {
		inventoryService inventoryService.InventoryService
		config           *config.Config
	}
)

func NewInventoryQueueHandler(inventoryService inventoryService.InventoryService, config *config.Config) InventoryQueueHandler {
	return &inventoryQueueHandler{inventoryService, config}
}
