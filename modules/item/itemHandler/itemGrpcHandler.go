package itemHandler

import (
	"context"

	itemPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemService"
)

type (
	itemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemService itemService.ItemService
	}
)

func NewItemGrpcHandler(itemService itemService.ItemService) *itemGrpcHandler {
	return &itemGrpcHandler{
		UnimplementedItemGrpcServiceServer: itemPb.UnimplementedItemGrpcServiceServer{},
		itemService:                        itemService,
	}
}

func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}
