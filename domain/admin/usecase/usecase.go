package usecase

import (
	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
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

	return
}

func (u *usecase) Login(req request.Login) (data interface{}, err error) {

	return
}
