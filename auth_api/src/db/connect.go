package database

import (
	"fmt"

	"github.com/renasami/svelte-go/auth_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	schema         = "%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	username       = "mariadb"
	password       = "secret"
	dbName         = "auth"
	datasourceName = fmt.Sprintf(schema, username, password, dbName)
	// DBインスタンス
	DB *gorm.DB
)

func Connect() {
	connection, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	// コネクション情報を追加
	DB = connection
	connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
	connection.AutoMigrate(&models.User{})
}
