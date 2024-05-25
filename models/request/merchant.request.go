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
	MerchantId       *string `query:"merchantId"`
	Limit            *int    `query:"limit"`
	Offset           *int    `query:"offset"`
	Name             *string `query:"name"`
	MerchantCategory *string `query:"merchantCategory"`
	CreatedAt        *string `query:"createdAt"`
}
