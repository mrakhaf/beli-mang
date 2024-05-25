package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/models/response"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type repoHandler struct {
	databaseDB *sql.DB
}

func NewRepository(databaseDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		databaseDB: databaseDB,
	}
}

func (r *repoHandler) GetUserByEmailAndRole(email string, role string) (user entity.User, err error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' AND role = '%s'", email, role)

	row := r.databaseDB.QueryRow(query)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt)

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (r *repoHandler) SaveUser(req request.Register) (user entity.User, err error) {

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return
	}

	user = entity.User{
		ID:        utils.GenerateUUID(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  password,
		Role:      "admin",
		CreatedAt: time.Now().Format(time.RFC3339Nano),
	}

	query := fmt.Sprintf("INSERT INTO users (id, username, email, password, role, created_at) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", user.ID, user.Username, user.Email, user.Password, user.Role, user.CreatedAt)

	_, err = r.databaseDB.Exec(query)
	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetUserByUsername(username string) (user entity.User, err error) {

	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)

	row := r.databaseDB.QueryRow(query)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt)

	return
}

func (r *repoHandler) SaveMerchant(req request.MerchantRequest) (merchant entity.Merchant, err error) {

	merchant = entity.Merchant{
		ID:               utils.GenerateUUID(),
		Name:             req.Name,
		MerchantCategory: req.MerchantCategory,
		ImageUrl:         req.ImageUrl,
		Latitude:         req.Location.Lat,
		Longitude:        req.Location.Long,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}

	query := fmt.Sprintf("INSERT INTO merchant (id, name, merchantcategory, imageurl, latitude, longitude, created_at) VALUES ('%s', '%s', '%s', '%s', '%f', '%f', '%s')", merchant.ID, merchant.Name, merchant.MerchantCategory, merchant.ImageUrl, merchant.Latitude, merchant.Longitude, merchant.CreatedAt)

	_, err = r.databaseDB.Exec(query)
	if err != nil {
		return
	}

	return

}

func (r *repoHandler) GetMerchants(req request.GetMerchants) (merchants []entity.Merchant, meta response.Meta, err error) {

	query := "SELECT * FROM merchant WHERE 1 = 1"
	queryTotal := "SELECT COUNT(*) FROM merchant WHERE 1 = 1"

	if req.MerchantId != nil {
		query += fmt.Sprintf(" AND id = '%s'", *req.MerchantId)
		queryTotal += fmt.Sprintf(" AND id = '%s'", *req.MerchantId)
	}

	if req.Name != nil {
		query += fmt.Sprintf(" AND name LIKE '%%%s%%'", *req.Name)
		queryTotal += fmt.Sprintf(" AND name LIKE '%%%s%%'", *req.Name)
	}

	if req.MerchantCategory != nil {
		query += fmt.Sprintf(" AND merchantcategory = '%s'", *req.MerchantCategory)
		queryTotal += fmt.Sprintf(" AND merchantcategory = '%s'", *req.MerchantCategory)
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if *req.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	if req.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *req.Limit)
		meta.Limit = *req.Limit
	} else {
		query += fmt.Sprintf(" LIMIT 5")
		meta.Limit = 5
	}

	if req.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *req.Offset)
		meta.Offset = *req.Offset
	} else {
		query += fmt.Sprintf(" OFFSET 0")
		meta.Offset = 0
	}

	//process query item
	fmt.Println(query)
	rows, err := r.databaseDB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var merchant entity.Merchant
		err = rows.Scan(&merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageUrl, &merchant.Latitude, &merchant.Longitude, &merchant.CreatedAt)
		if err != nil {
			return
		}
		merchants = append(merchants, merchant)
	}

	//process query total
	rowTotal := r.databaseDB.QueryRow(queryTotal)
	err = rowTotal.Scan(&meta.Total)
	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetMerchantsById(merchantId string) (merchant entity.Merchant, err error) {

	query := fmt.Sprintf("SELECT * FROM merchant WHERE id = '%s'", merchantId)

	row := r.databaseDB.QueryRow(query)

	err = row.Scan(&merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageUrl, &merchant.Latitude, &merchant.Longitude, &merchant.CreatedAt)

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	return
}

func (r *repoHandler) SaveItem(req request.CreateItem, merchantId string) (item entity.Item, err error) {

	item = entity.Item{
		ID:              utils.GenerateUUID(),
		MerchantID:      merchantId,
		Name:            req.Name,
		ProductCategory: req.ProductCategory,
		Price:           req.Price,
		ImageUrl:        req.ImageUrl,
		CreatedAt:       time.Now().Format(time.RFC3339Nano),
	}

	query := fmt.Sprintf("INSERT INTO merchant_item (id, merchant_id, name, productcategory, price, imageurl, created_at) VALUES ('%s', '%s', '%s', '%s', '%d', '%s', '%s')", item.ID, item.MerchantID, item.Name, item.ProductCategory, item.Price, item.ImageUrl, item.CreatedAt)

	_, err = r.databaseDB.Exec(query)
	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetItems(req request.GetItems, merchantId string) (items []entity.Item, meta response.Meta, err error) {
	query := fmt.Sprintf("SELECT * FROM merchant_item WHERE merchant_id = '%s'", merchantId)
	queryMeta := fmt.Sprintf("SELECT COUNT(*) FROM merchant_item WHERE merchant_id = '%s'", merchantId)

	if req.Name != nil {
		query += fmt.Sprintf(" AND name LIKE '%%%s%%'", *req.Name)
		queryMeta += fmt.Sprintf(" AND name LIKE '%%%s%%'", *req.Name)
	}

	if req.ProductCategory != nil {
		query += fmt.Sprintf(" AND productcategory = '%s'", *req.ProductCategory)
		queryMeta += fmt.Sprintf(" AND productcategory = '%s'", *req.ProductCategory)
	}

	if req.ItemId != nil {
		query += fmt.Sprintf(" AND id = '%s'", *req.ItemId)
		queryMeta += fmt.Sprintf(" AND id = '%s'", *req.ItemId)
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if *req.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	if req.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *req.Limit)
		meta.Limit = *req.Limit
	} else {
		query += fmt.Sprintf(" LIMIT 5")
		meta.Limit = 5
	}

	if req.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *req.Offset)
		meta.Offset = *req.Offset
	} else {
		query += fmt.Sprintf(" OFFSET 0")
		meta.Offset = 0
	}

	//process query item
	fmt.Println(query)
	rows, err := r.databaseDB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.Item
		err = rows.Scan(&item.ID, &item.MerchantID, &item.Name, &item.ProductCategory, &item.Price, &item.ImageUrl, &item.CreatedAt)
		if err != nil {
			return
		}
		items = append(items, item)
	}

	//process query total
	rowTotal := r.databaseDB.QueryRow(queryMeta)
	err = rowTotal.Scan(&meta.Total)
	if err != nil {
		return
	}

	return
}
