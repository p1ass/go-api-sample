package persistance

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewSqlDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		"plus",
		"plus",
		"db",
		"3306",
		"go-api",
	)

	conn, err := gorm.Open("mysql", connectionString)

	if nil != err {
		panic(err)
	}

	err = conn.DB().Ping()

	if nil != err {
		panic(err)
	}
	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
