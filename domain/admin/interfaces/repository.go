package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/models/response"
)

type Repository interface {
	GetUserByEmailAndRole(email string, role string) (user entity.User, err error)
	SaveUser(req request.Register) (data entity.User, err error)
	GetUserByUsername(username string) (user entity.User, err error)
	SaveMerchant(req request.MerchantRequest) (merchant entity.Merchant, err error)
	GetMerchants(req request.GetMerchants) (merchants []entity.Merchant, meta response.Meta, err error)
	GetMerchantsById(merchantId string) (merchant entity.Merchant, err error)
	SaveItem(req request.CreateItem, merchantId string) (item entity.Item, err error)
	GetItems(req request.GetItems, merchantId string) (items []entity.Item, meta response.Meta, err error)
}
