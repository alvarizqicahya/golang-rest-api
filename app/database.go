package app

import (
	"log"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var InitDB *gorm.DB

func DatabaseConnect() {
	var err error
	InitDB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test?parseTime=True&loc=Asia%2FJakarta&charset=utf8"), &gorm.Config{})
	if err != nil {
		log.Panic("Koneksi database error, " + err.Error())
	}

	log.Println("Koneksi database berhasil")
}

func DatabaseClose() {
	dbConn, _ := InitDB.DB()

	dbConn.Close()

	log.Println("Koneksi database tertutup")
}
