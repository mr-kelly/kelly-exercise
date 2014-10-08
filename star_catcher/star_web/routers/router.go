package routers

import (
	"github.com/mr-kelly/beego"
	"star_web/controllers"
)

func init() {
	var defaultController = &controllers.MainController{}

	beego.Router("/", defaultController, "Get:GetIndex")
	beego.Router("/star/:id", defaultController, "Get:GetStar")
}
