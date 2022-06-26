package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/macduyhai/BaseBE/config"
	"github.com/macduyhai/BaseBE/rounters"
)

func main() {
	log.Println("Start backend basic !")
	// Load config from env file
	conf := config.NewConfig()
	if conf == nil {
		// log.Prinln("Read config error")
		// return
		log.Fatal("Read config error")
	} else {
		// log.Println(conf)
	}
	// use gorm to init conection DB Mysql
	db, err := gorm.Open("mysql", conf.MYSQLURL)
	if err != nil {
		panic("Open DB error:" + err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Close DB error")
		}
	}()
	// Init router request
	router := rounters.NewRounter(conf, db)
	app, err := router.InitGin()
	if err != nil {
		log.Println("Init Gin error:" + err.Error())
		return
	}
	// Run APP port 88
	log.Println("Run App")
	_ = app.Run(conf.PortEngine)

}
