package persistance

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewSQLDB creates a connection to MySQL server.
func NewSQLDB() *gorm.DB {

	// In production environment, you must read config from environment variables.
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		"user",
		"password",
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
