package interfaces

import "github.com/mrakhaf/halo-suster/models/request"

type Usecase interface {
	Register(req request.Register) (data interface{}, err error)
	Login(req request.Login) (data interface{}, err error)
	CreateMerchant(req request.MerchantRequest) (data interface{}, err error)
	GetMerchants(req request.GetMerchants) (data interface{}, err error)
	CreateItem(req request.CreateItem, merchantId string) (data interface{}, err error)
	GetItems(req request.GetItems, merchantId string) (data interface{}, err error)
}
