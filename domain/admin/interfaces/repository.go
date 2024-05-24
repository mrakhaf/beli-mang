package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	GetUserByEmailAndRole(email string, role string) (user entity.User, err error)
	SaveUser(req request.Register) (data entity.User, err error)
	GetUserByUsername(username string) (user entity.User, err error)
}
