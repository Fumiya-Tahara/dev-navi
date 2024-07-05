package gorm

import (
	"github.com/Fumiya-Tahara/dev-navi/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur UserRepository) Create(u *entity.User) (*entity.User, error) {
	err := ur.DB.Create(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur UserRepository) FindByID(id uint) (*entity.User, error) {
	user := &entity.User{}

	err := ur.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Update(u *entity.User) (*entity.User, error) {
	user := &entity.User{}

	err := ur.DB.First(user, u.ID).Error
	if err != nil {
		return nil, err
	}

	user.Name = u.Name

	err = ur.DB.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Delete(id uint) error {
	err := ur.DB.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
