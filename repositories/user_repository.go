package repositories

import (
	"ex2/models"
	"ex2/storages"
)

type userRepository struct {
	db *storages.PSQLManager
}
type UserRepository interface {
	ReadTokenByToken(token string) (*models.AuthToken, error)
	ReadUserByUserId(userId int64) (*models.User, error)
}

func NewUserRepository(db *storages.PSQLManager) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (u *userRepository) ReadTokenByToken(token string) (*models.AuthToken, error) {
	authen := models.AuthToken{}
	if err := u.db.Where(&models.AuthToken{Token: token}).First(&authen).Error; err != nil {
		return nil, err
	}
	return &authen, nil
}
func (u *userRepository) ReadUserByUserId(userId int64) (*models.User, error) {
	user := models.User{}
	if err := u.db.Where(&models.User{Id: userId}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
