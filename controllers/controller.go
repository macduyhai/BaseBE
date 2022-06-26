package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/macduyhai/BaseBE/dtos"
	"github.com/macduyhai/BaseBE/models"
	"github.com/macduyhai/BaseBE/services"
	"github.com/macduyhai/BaseBE/utilitys"
)

type Controller struct {
	userService services.UserService
}

func NewController(provider services.Provider) Controller {
	return Controller{
		userService: provider.GetUserService(),
	}
}
func (ctl *Controller) Login(context *gin.Context) {
	// log.Println(context.Request.Header)
	var request dtos.LoginRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json request error: " + err.Error())
		utilitys.ResponseError400(context, "Login error")
		return
	}

	token, err := ctl.userService.Login(request)
	if err != nil {
		log.Println("Create token login Error:" + err.Error())
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, token, "Login success")
	}
}

func (ctl *Controller) Create(context *gin.Context) {
	var request dtos.AddRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json request error: " + err.Error())
		utilitys.ResponseError400(context, err.Error())
		return
	}
	timeNow := utilitys.TimeIn("Asia/Ho_Chi_Minh")
	staff := models.Staff{
		Username:   request.Username,
		Password:   request.Password,
		Fullname:   request.Fullname,
		Salary:     request.Salary,
		CreateTime: &timeNow,
	}
	user, err := ctl.userService.Create(staff)
	if err != nil {
		log.Println("Create user error:" + err.Error())
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, user, "Create success")
	}
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong Pong",
	})
}
func stringYYYYMMDD2Time(input string) (*time.Time, error) {
	layout := "2006-01-02"
	result, err := time.Parse(layout, input)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
