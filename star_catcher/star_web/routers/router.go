package routers

import (
	"github.com/mr-kelly/beego"
	"star_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
