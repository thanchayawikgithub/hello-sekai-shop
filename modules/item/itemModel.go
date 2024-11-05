package item

import "github.com/thanchayawikgithub/hello-sekai-shop/pkg/models"

type (
	CreateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageURL string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	ItemShowCase struct {
		ItemID   string  `json:"item_id"`
		Title    string  `json:"title"`
		ImageURL string  `json:"image_url"`
		Price    float64 `json:"price"`
		Damage   int     `json:"damage"`
	}

	ItemSearchReq struct {
		Title string `json:"title" validate:"required,max=64"`
		models.PaginateReq
	}

	ItemUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageURL string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	EnableOrDisableItemReq struct {
		Status bool `json:"status"`
	}
)
