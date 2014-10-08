package models

import (
	"github.com/mr-kelly/beego/orm"
)

type Star struct {
	Id   int
	Name string
}

type Movie struct {
	Id     string // fanhao
	StarId int
}

func init() {

	orm.RegisterModel(new(Star))
	orm.RegisterModel(new(Movie))
}
