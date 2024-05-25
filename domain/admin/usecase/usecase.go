package usecase

import (
	"fmt"

	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/models/response"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type (
	usecase struct {
		repository interfaces.Repository
		JwtAccess  *jwt.JWT
	}
)

func NewUsecase(repository interfaces.Repository, JwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		JwtAccess:  JwtAccess,
	}
}

func (u *usecase) Register(req request.Register) (data interface{}, err error) {
	//check user by email
	user, err := u.repository.GetUserByEmailAndRole(req.Email, "admin")
	if err != nil {
		return
	}

	if user != (entity.User{}) {
		err = fmt.Errorf("user already exist")
		return
	}

	//save user
	user, err = u.repository.SaveUser(req)
	if err != nil {
		return
	}

	//generate token
	token, err := u.JwtAccess.GenerateToken(user.Username, "admin")
	if err != nil {
		return
	}

	data = map[string]interface{}{
		"token": token,
	}

	return
}

func (u *usecase) Login(req request.Login) (data interface{}, err error) {

	user, err := u.repository.GetUserByUsername(req.Username)
	if err != nil {
		return
	}

	if user.Role != "admin" {
		err = fmt.Errorf("user not admin")
		return
	}

	if err = utils.CheckPasswordHash(req.Password, user.Password); err != nil {
		return
	}

	//generate token
	token, err := u.JwtAccess.GenerateToken(user.Username, "admin")
	if err != nil {
		return
	}

	data = map[string]interface{}{
		"token": token,
	}

	return
}

func (u *usecase) CreateMerchant(req request.MerchantRequest) (data interface{}, err error) {

	merchant, err := u.repository.SaveMerchant(req)

	if err != nil {
		return
	}

	data = map[string]interface{}{
		"merchantId": merchant.ID,
	}

	return

}

func (u *usecase) GetMerchants(req request.GetMerchants) (data interface{}, err error) {

	merchants, meta, err := u.repository.GetMerchants(req)

	if err != nil {
		return
	}

	var dataMerchant []response.Merchants

	if len(merchants) > 0 {
		dataMerchant = []response.Merchants{}

		for _, merchant := range merchants {

			dataMerchant = append(dataMerchant, response.Merchants{
				MerchantId:       merchant.ID,
				Name:             merchant.Name,
				MerchantCategory: merchant.MerchantCategory,
				ImageUrl:         merchant.ImageUrl,
				Location: response.Location{
					Lat:  merchant.Latitude,
					Long: merchant.Longitude,
				},
				CreatedAt: merchant.CreatedAt,
			})
		}
	} else {
		dataMerchant = []response.Merchants{}
	}

	data = map[string]interface{}{
		"data": dataMerchant,
		"meta": meta,
	}

	return
}
