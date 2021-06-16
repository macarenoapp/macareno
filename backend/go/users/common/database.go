package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//_ "github.com/go-sql-driver/mysql"
// _ "github.com/lib/pq"

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("mysql",  DbURL(BuildDBConfig())) //gorm.Open(mysql) from mysql //(postgres) for postgres
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}
type DBConfig struct {
    Host     string
    Port     int
    User     string
    DBName   string
    Password string
}




func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "mysql",
		Port:     3306,  //3306 for mysql  //5432 for postgres
		User:     "xema", //xema for mysql  //postgres
		Password: "123456789",
		DBName:   "gran_melon", //first_go for mysql  //gran_melon for postgres
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	// return psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    // "password=%s dbname=%s sslmode=disable",
    // host, port, user, password, dbname)
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
