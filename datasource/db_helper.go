package datasource

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDbInstance() *gorm.DB {
	db, err := gorm.Open(beego.AppConfig.String("mysqldriver"), fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqldb"),
	))
	if err != nil {
		log.Fatalf("mysql connect error:%v", err)
	}
	db.SingularTable(true)
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置数据库连接最大打开数。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置可重用连接的最长时间
	db.DB().SetConnMaxLifetime(time.Hour)

	db.LogMode(true)
	return db
}
