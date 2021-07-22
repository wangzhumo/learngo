package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/models"
)

type BlackIpService interface {
	Get(id int) *models.IpBlack
	GetAll(page, size int) []models.IpBlack
	CountAll() int64
	Search(ip string) []models.IpBlack
	Update(user *models.IpBlack, columns []string) error
	Insert(user *models.IpBlack) error
	GetByIp(ip string) *models.IpBlack
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func (b *blackIpService) Get(id int) (*models.IpBlack, error) {
	return b.dao.Get(id)
}

func (b *blackIpService) GetAll(page, size int) ([]models.IpBlack, error) {
	return b.dao.GetAll(page, size)
}

func (b *blackIpService) CountAll() int64 {
	return b.dao.CountAll()
}

func (b *blackIpService) Search(ip string) []models.IpBlack {
	return b.dao.Search(ip)
}

func (b *blackIpService) Update(ip *models.IpBlack, columns []string) error {
	return b.dao.Update(ip, columns)
}

func (b *blackIpService) Insert(ip *models.IpBlack) error {
	return b.dao.Insert(ip)
}

func (b *blackIpService) GetByIp(ip string) *models.IpBlack {
	return b.dao.GetByIp(ip)
}
