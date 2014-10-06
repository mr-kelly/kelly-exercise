package models

import (
	"github.com/astaxie/beego/orm"
)

type Star struct {
	Id int

}

func init() {
	
	orm.RegisterModel(new(Star))

}