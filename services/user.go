package services

import (
	"encoding/base64"
	"encoding/binary"
	"log"
	_ "log"

	"github.com/macduyhai/BaseBE/daos"
	"github.com/macduyhai/BaseBE/models"

	"github.com/macduyhai/BaseBE/config"
	"github.com/macduyhai/BaseBE/dtos"
	"github.com/macduyhai/BaseBE/middlewares"
)

type UserService interface {
	Login(request dtos.LoginRequest) (*dtos.LoginResponse, error)
	Create(user models.Staff) (*models.Staff, error)
	Edit(user models.Staff, token string) (*models.Staff, error)
	Delete(user models.Staff, token string) error
}
type userServiceImpl struct {
	config  *config.Config
	userDao daos.UserDao
	jwt     middlewares.JWT
}

func NewUserService(conf *config.Config, userDao daos.UserDao, jwt middlewares.JWT) UserService {
	return &userServiceImpl{
		config:  conf,
		userDao: userDao,
		jwt:     jwt,
	}
}

func (service *userServiceImpl) Login(request dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := service.userDao.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	// Encode userID by RSA ->token
	byteID := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteID, uint64(user.ID))
	encodeID, err := RsaEncrypt(byteID, []byte(service.config.PublicKey))
	if err != nil {
		return nil, err

	}

	tokenString := base64.StdEncoding.EncodeToString(encodeID)
	token, err := service.jwt.CreateTokenPrivate(tokenString)
	if err != nil {
		return nil, err

	}
	response := dtos.LoginResponse{
		Token: token,
	}
	return &response, nil

}
func (service *userServiceImpl) Create(user models.Staff) (*models.Staff, error) {
	return service.userDao.Create(user)

}
func (service *userServiceImpl) Edit(user models.Staff, token string) (*models.Staff, error) {
	id, err := service.DecodeToken(token)
	if err != nil {
		log.Println("Decode token edit error:" + err.Error())
		return nil, err
	}
	return service.userDao.Edit(user, id)
}
func (service *userServiceImpl) Delete(user models.Staff, token string) error {
	id, err := service.DecodeToken(token)
	if err != nil {
		log.Println("Decode token edit error:" + err.Error())
		return err
	}
	return service.userDao.Delete(user, id)
}

// func (service *userServiceImpl) FindMyID(id int64) *models.Staff {
// 	return service.userDao.FindMyID(id)
// }
