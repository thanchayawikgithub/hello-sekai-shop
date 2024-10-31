package itemHandler

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemService"

type (
	itemGrpcHandler struct {
		itemService itemService.ItemService
	}
)

func NewItemGrpcHandler(itemService itemService.ItemService) *itemGrpcHandler {
	return &itemGrpcHandler{itemService}
}
