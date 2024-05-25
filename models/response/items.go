package response

type Items struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	Price           int    `json:"price"`
	ImageUrl        string `json:"imageUrl"`
	CreatedAt       string `json:"createdAt"`
}

type ItemsResponse struct {
	Data []Items `json:"data"`
	Meta Meta    `json:"meta"`
}
