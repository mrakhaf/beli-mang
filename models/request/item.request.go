package request

type CreateItem struct {
	Name            string `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string `json:"productCategory" validate:"required,oneof=Food Beverage Snack Condiments Additions"`
	Price           int    `json:"price" validate:"required"`
	ImageUrl        string `json:"imageUrl" validate:"required,url"`
}

type GetItems struct {
	ItemId          *string `query:"itemId" validate:"omitempty"`
	Limit           *int    `query:"limit" validate:"omitempty"`
	Offset          *int    `query:"offset" validate:"omitempty"`
	Name            *string `query:"name" validate:"omitempty"`
	ProductCategory *string `query:"productCategory" validate:"omitempty"`
	CreatedAt       *string `query:"createdAt" validate:"omitempty"`
}
