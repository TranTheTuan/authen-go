package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/TranTheTuan/authen-go/app/domain/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindUserByUsername(username string) (*model.User, error) {
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

func (repo *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := repo.db.Transaction(func(db *gorm.DB) error {
		if err := db.Create(&user).Error; err != nil {
			return err
		}

		user.CasbinUser = fmt.Sprint(user.ID)
		if err := db.Model(model.User{}).Save(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
}
