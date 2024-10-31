package itemHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemService"
)

type (
	ItemHttpHandler interface{}

	itemHttpHandler struct {
		itemService itemService.ItemService
		config      *config.Config
	}
)

func NewItemHttpHandler(config *config.Config, itemService itemService.ItemService) ItemHttpHandler {
	return &itemHttpHandler{itemService, config}
}
