package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	db, err := orm.GetDB()

	if err != nil {
	    fmt.Println("get default DataBase")
	    fmt.Println(db)
	}



	this.TplNames = "index.tpl"
}
