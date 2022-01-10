package database

import (
	"github.com/klishchov-bohdan/myrepos/internal/models"
	"github.com/klishchov-bohdan/myrepos/pkg/dbconn"
)

type UserDBRepository struct {
	db dbconn.DBConn
}

func NewUserDBRepository(db dbconn.DBConn) *UserDBRepository {
	return &UserDBRepository{
		db: db,
	}
}

func (u *UserDBRepository) GetByEmail(email string) (user *models.User, err error) {
	// SELECT email, password_hash, created_at, FROM users WHERE email = email
	res := u.db.GetDB().Where("email = ?", email).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return
}

func (u *UserDBRepository) Create(user *models.User) (*models.User, error) {
	res := u.db.GetDB().Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (u *UserDBRepository) GetByID(id int) (user *models.User, err error) {
	res := u.db.GetDB().First(user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return
}
