package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/CienciaArgentina/roles/internal/role"
	_ "github.com/go-sql-driver/mysql"
)

// Build Builds and returns productive App
func Build() *App {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Print(err)
		panic("Couldn't creating connection to database")
	}
	if err = db.Ping(); err != nil {
		fmt.Print(err)
		panic("Error pinging database")
	}

	roleDAO := role.NewDAO(db)
	roleService := role.NewService(roleDAO)
	roleController := role.NewController(roleService)
	return New(roleController)
}
