package daos

import (
	_ "log"

	"github.com/macduyhai/BaseBE/models"

	"github.com/jinzhu/gorm"
)

type UserDao interface {
	Login(username, pass string) (*models.Staff, error)
	Create(user models.Staff) (*models.Staff, error)
}

type userDaoImpl struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDaoImpl{db: db}
}

func (dao *userDaoImpl) Login(username, pass string) (*models.Staff, error) {
	var user models.Staff
	if err := dao.db.Where("username = ? AND password = ?", username, pass).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
func (dao *userDaoImpl) Create(user models.Staff) (*models.Staff, error) {

	if err := dao.db.Create(&user).Error; err != nil {
		return nil, err

	}
	return &user, nil
}
