package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}
