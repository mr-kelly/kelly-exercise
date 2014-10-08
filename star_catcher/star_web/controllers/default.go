package controllers

import (
	"fmt"
	"github.com/mr-kelly/beego"
	"github.com/mr-kelly/beego/orm"
	"star_web/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) GetIndex() {
	o := orm.NewOrm()
	// o, err := orm.GetDB()

	// if err != nil {
	// 	fmt.Println("get default DataBase")
	// 	fmt.Println(o)
	// }
	var users []*models.Star
	num, err := o.QueryTable("star").All(&users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	this.Data["Stars"] = users
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

func (this *MainController) GetStar() {
	star_id := this.Ctx.Input.Param(":id")

	this.Data["Star"] = star_id
	this.Layout = "layout.tpl"
	this.TplNames = "star_view.tpl"
}