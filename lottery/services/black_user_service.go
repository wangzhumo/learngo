package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/models"
)

type UserService interface {
	Get(uid int) (*models.UserBlack, error)
	GetAll(page int, size int) ([]models.UserBlack, error)
	CountAll() int64
	Update(user *models.UserBlack, columns []string) error
	Insert(user *models.UserBlack) error
}

type blackUserService struct {
	dao *dao.BlackUserDao
}

func (b *blackUserService) Get(uid int) (*models.UserBlack, error) {
	return b.dao.Get(uid)
}

func (b *blackUserService) GetAll(page int, size int) ([]models.UserBlack, error) {
	return b.dao.GetAll(page, size)
}

func (b *blackUserService) CountAll() int64 {
	return b.dao.CountAll()
}

func (b *blackUserService) Update(user *models.UserBlack, columns []string) error {
	return b.dao.Update(user, columns)
}

func (b *blackUserService) Insert(user *models.UserBlack) error {
	return b.dao.Insert(user)
}
