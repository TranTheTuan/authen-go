package repository

import (
	"authen-go/app/domain/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Login(username string) (*model.User, error) {
	var tmpUser model.User
	err := repo.db.Where("username = ?", username).First(&tmpUser).Error
	if err != nil {
		return nil, err
	}
	if tmpUser.ID > 0 {
		return &tmpUser, nil
	}
	return nil, err
}

func (repo *UserRepository) IsUsernameExist(username string) (bool, error) {
	var tmpObj model.User
	err := repo.db.Where("username = ?", username).First(&tmpObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tmpObj.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (repo *UserRepository) Register(user *model.User) (*model.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
