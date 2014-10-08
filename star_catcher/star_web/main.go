package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mr-kelly/beego"
	"github.com/mr-kelly/beego/orm"
	_ "star_web/routers"
	// "star_web/models"
)

func main() {

	orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "./data.db")

	beego.Run()
}
