package mysql_conn

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type MySQLConn struct {
	DB *gorm.DB
}

func NewMySQLConn() (*MySQLConn, error) {
	err := godotenv.Load("config/mysql.env")
	if err != nil {
		return nil, err
	}
	dbuser := os.Getenv("user")
	dbpwd := os.Getenv("pwd")
	dbname := os.Getenv("dbName")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpwd, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQLConn{
		DB: db,
	}, nil
}

func (conn *MySQLConn) SetConnPool(maxConns, idleConns int, lifeTimeConns time.Duration) error {
	db, err := conn.DB.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(idleConns)
	db.SetConnMaxLifetime(lifeTimeConns)
	return nil
}

func (conn *MySQLConn) GetDB() *gorm.DB {
	return conn.DB
}
