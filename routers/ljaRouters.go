package routers

import (
	"FabricWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/blade-goods/goods-batch/input-info", &controllers.GoodsBatchController{}, "Get:InputInfo")

	beego.Router("/blade-goods/goods/submit", &controllers.GoodsController{}, "Get:SubmitGet;Post:SubmitPost")

	beego.Router("/blade-goods/goods/update", &controllers.GoodsController{}, "Get:UpdateGet;Post:UpdatePost")

	beego.Router("/blade-goods/goods/detail", &controllers.GoodsController{}, "Get:DetailGet;Post:DetailPost")
}
