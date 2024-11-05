package inventoryHandler

import (
	"context"

	inventoryPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"
)

type (
	inventoryGrpcHandler struct {
		inventoryPb.UnimplementedInventoryGrpcServiceServer
		inventoryService inventoryService.InventoryService
	}
)

func NewInventoryGrpcHandler(inventoryService inventoryService.InventoryService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		UnimplementedInventoryGrpcServiceServer: inventoryPb.UnimplementedInventoryGrpcServiceServer{},
		inventoryService:                        inventoryService,
	}
}

func (g *inventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellReq) (*inventoryPb.IsAvailableToSellRes, error) {
	return nil, nil
}
