package models

import (
	"fmt"
	"gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err         error
		dbTaye      string
		dbName      string
		user        string
		password    string
		host        string
		tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("读取database配置错误：%v", err)
	}
	dbTaye = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbTaye, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetConnMaxIdleTime(10)
	db.DB().SetMaxIdleConns(100)
}
func Close() {
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("数据库关闭失败：%v", err)
		}
	}(db)
}
