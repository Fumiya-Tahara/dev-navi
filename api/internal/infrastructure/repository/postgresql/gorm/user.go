package gorm

import (
	"github.com/Fumiya-Tahara/dev-navi/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur UserRepository) Create(u *entity.User) (user *entity.User, err error) {
	err = ur.DB.Create(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (UserRepository) FindById(id uint) (*entity.User, error) {
	return nil, nil
}

func (UserRepository) Update(user *entity.User) error {
	return nil
}

func (UserRepository) Delete(user *entity.User, id uint) error {
	return nil
}
