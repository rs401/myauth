package repository

import (
	"errors"

	"github.com/rs401/myauth/auth/models"
	"github.com/rs401/myauth/db"
	"gorm.io/gorm"
)

var ErrorBadID error = errors.New("bad id")

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id uint) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Update(user *models.User) error
	Delete(id uint) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{db: conn.DB()}
}

func (r *usersRepository) Save(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *usersRepository) GetById(id uint) (user *models.User, err error) {
	result := r.db.Where("ID = ?", id).First(&user)
	return user, result.Error
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	result := r.db.Where("email = ?", email).Find(&user)
	return user, result.Error
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *usersRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *usersRepository) Delete(id uint) error {
	var user models.User
	r.db.Find(&user, id)
	if user.ID == 0 {
		return ErrorBadID
	}
	return r.db.Delete(&user).Error
}

func (r *usersRepository) DeleteAll() error {
	// return r.db.Where("1 = 1").Delete(&models.User{}).Error
	return r.db.Exec("DELETE FROM users").Error
}
