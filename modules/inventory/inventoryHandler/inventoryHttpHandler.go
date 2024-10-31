package inventoryHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"
)

type (
	InventoryHttpHandler interface{}

	inventoryHttpHandler struct {
		inventoryService inventoryService.InventoryService
		config           *config.Config
	}
)

func NewInventoryHttpHandler(config *config.Config, inventoryService inventoryService.InventoryService) InventoryHttpHandler {
	return &inventoryHttpHandler{inventoryService, config}
}
