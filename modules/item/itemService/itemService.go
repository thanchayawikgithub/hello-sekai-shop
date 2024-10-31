package itemService

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemRepository"
)

type (
	ItemService interface{}

	itemService struct {
		itemRepo itemRepository.ItemRepository
	}
)

func NewItemService(itemRepo itemRepository.ItemRepository) ItemService {
	return &itemService{itemRepo}
}
