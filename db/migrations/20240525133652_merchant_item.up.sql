CREATE TABLE merchant_item (
    id VARCHAR(255) PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    productcategory VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    imageurl VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)