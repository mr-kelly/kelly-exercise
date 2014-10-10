package controllers

import (
	"fmt"
	"github.com/mr-kelly/beego"
	"github.com/mr-kelly/beego/orm"
	"star_web/models"
	"strconv"
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

func (this *MainController) GetTemplate() {
	this.TplNames = "template.tpl"
}

func (this *MainController) GetStar() {
	star_id_param := this.Ctx.Input.Param(":id")
	star_id, err := strconv.Atoi(star_id_param)

	o := orm.NewOrm()

	star := models.Star{Id: star_id}

	err = o.Read(&star)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(star.Id, star.Name)
	}

	this.Data["Star"] = star
	this.Layout = "layout.tpl"
	this.TplNames = "star_view.tpl"
}
