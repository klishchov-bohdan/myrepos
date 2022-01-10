package main

import (
	"github.com/klishchov-bohdan/myrepos/internal/repositories/database"
	"github.com/klishchov-bohdan/myrepos/internal/routes"
	"github.com/klishchov-bohdan/myrepos/internal/services"
	"github.com/klishchov-bohdan/myrepos/pkg/dbconn/mysql_conn"
	"log"
	"net/http"
	"time"
)

func main() {
	mysqlConn, err := mysql_conn.NewMySQLConn()
	if err != nil {
		log.Fatal(err)
	}
	err = mysqlConn.SetConnPool(10, 5, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	//repo := &filesystem.UserFileRepository{}
	dbrepo := database.NewUserDBRepository(mysqlConn)
	service := services.New(dbrepo)
	routes.SetRoutes(service)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
