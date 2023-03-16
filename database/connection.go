package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
// 	var connectionString = fmt.Sprintf(
// 		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		dbUser, dbPassword, dbName,
// 	)

// 	return gorm.Open("mysql", connectionString)
// }

func ConnectToDB(dbUser string, dbPassword string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		// USING SQLSERVER
		"sqlserver://%s:%s@(IPADDRESS)?database=(DBNAME)",
		dbUser, dbPassword,
	)

	return gorm.Open("mssql", connectionString)
}
