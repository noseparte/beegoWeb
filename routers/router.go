package routers

import (
	"beegoWeb/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gm", &controllers.GmController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/player", &controllers.PlayerController{})
}
