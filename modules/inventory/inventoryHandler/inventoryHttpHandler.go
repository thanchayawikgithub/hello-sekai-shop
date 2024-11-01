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

func NewInventoryHttpHandler(inventoryService inventoryService.InventoryService, config *config.Config) InventoryHttpHandler {
	return &inventoryHttpHandler{inventoryService, config}
}
