package request

type Location struct {
	Lat  float64 `json:"lat" validate:"required"`
	Long float64 `json:"long" validate:"required"`
}

type MerchantRequest struct {
	Name             string   `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory string   `json:"merchantCategory" validate:"required,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	ImageUrl         string   `json:"imageUrl" validate:"required,url"`
	Location         Location `json:"location" validate:"required"`
}

type GetMerchants struct {
	MerchantId       *string `query:"merchantId" validate:"omitempty"`
	Limit            *int    `query:"limit" validate:"omitempty"`
	Offset           *int    `query:"offset" validate:"omitempty"`
	Name             *string `query:"name" validate:"omitempty"`
	MerchantCategory *string `query:"merchantCategory" validate:"omitempty"`
	CreatedAt        *string `query:"createdAt" validate:"omitempty"`
}
