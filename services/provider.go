package services

import (
	_ "log"

	"github.com/jinzhu/gorm"
	"github.com/macduyhai/BaseBE/config"
	"github.com/macduyhai/BaseBE/daos"
	"github.com/macduyhai/BaseBE/middlewares"
)

type Provider interface {
	GetUserService() UserService
}

type providerImpl struct {
	config *config.Config
	db     *gorm.DB
}

func NewProviderService(conf *config.Config, db *gorm.DB) Provider {
	return &providerImpl{config: conf, db: db}
}
func (p *providerImpl) GetUserService() UserService {
	userDao := daos.NewUserDao(p.db)
	jwtClient := middlewares.NewJWT(p.config.SecretKey)

	return NewUserService(p.config, userDao, jwtClient)
}
