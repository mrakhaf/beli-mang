CREATE TABLE merchant (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    merchantcategory VARCHAR(255) NOT NULL,
    imageurl VARCHAR(255) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);