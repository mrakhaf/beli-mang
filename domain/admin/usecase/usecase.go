package usecase

import (
	"fmt"

	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
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

	return
}
