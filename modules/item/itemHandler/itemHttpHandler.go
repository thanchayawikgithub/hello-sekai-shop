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

func NewItemHttpHandler(itemService itemService.ItemService, config *config.Config) ItemHttpHandler {
	return &itemHttpHandler{itemService, config}
}
