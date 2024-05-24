package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
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
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
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
