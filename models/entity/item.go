package entity

// id VARCHAR(255) PRIMARY KEY,
// merchant_id VARCHAR(255) NOT NULL,
// name VARCHAR(255) NOT NULL,
// productcategory VARCHAR(255) NOT NULL,
// price FLOAT NOT NULL,
// imageurl VARCHAR(255) NOT NULL,
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

type Item struct {
	ID              string
	MerchantID      string
	Name            string
	ProductCategory string
	Price           int
	ImageUrl        string
	CreatedAt       string
}
