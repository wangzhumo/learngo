package services

import (
	"com.wangzhumo.lottery/dao"
	"com.wangzhumo.lottery/datasource"
	"com.wangzhumo.lottery/models"
	"fmt"
	"strconv"
	"time"
)

type UserDayService interface {
	Get(id int) (*models.CountRecord,error)
	GetAll(page, size int) ([]models.CountRecord, error)
	Count(uid, day int) int
	CountAll() int64
	Search(uid, day int) []models.CountRecord
	Update(user *models.CountRecord, columns []string) error
	Insert(user *models.CountRecord) error
	GetUserToday(uid int) *models.CountRecord
}

type userDayService struct {
	dao *dao.CountDao
}

func (u *userDayService) Get(id int) (*models.CountRecord, error) {
	return u.dao.Get(id)
}

func (u *userDayService) GetAll(page int, size int) ([]models.CountRecord, error) {
	return u.dao.GetAll(page, size)
}

func (u *userDayService) Count(uid, day int) int {
	return u.dao.Count(uid, day)
}

func (u *userDayService) CountAll() int64 {
	return u.dao.CountAll()
}

func (u *userDayService) Search(uid, day int) []models.CountRecord {
	return u.dao.Search(uid, day)
}

func (u *userDayService) Update(user *models.CountRecord, columns []string) error {
	return u.dao.Update(user, columns)
}

func (u *userDayService) Insert(user *models.CountRecord) error {
	return u.dao.Insert(user)
}

func (u *userDayService) GetUserToday(uid int) *models.CountRecord {
	year, month, day := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", year, month, day)
	date, _ := strconv.Atoi(strDay)
	countRecords := u.dao.Search(uid, date)
	if countRecords != nil && len(countRecords) > 0 {
		return &countRecords[0]
	}
	return nil
}

func NewUserDayService() UserDayService {
	return &userDayService{
		dao: dao.NewCountDao(datasource.InstanceDB()),
	}
}