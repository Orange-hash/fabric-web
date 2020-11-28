package routers

import (
	"FabricWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/blade-goods/goods-batch/input-info", &controllers.GoodsBatchController{}, "Get:InputInfo")
}
