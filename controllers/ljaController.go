package controllers

import "github.com/astaxie/beego"

type GoodsBatchController struct {
	beego.Controller
}

type GoodsController struct {
	beego.Controller
}

func (c *GoodsBatchController) InputInfo() {
	//扫码回调方法
	//Get请求
}

func (c *GoodsController) SubmitGet() {
	//商品新增方法
	//Get请求
	//渲染 新增商品页面
}

func (c *GoodsController) SubmitPost() {
	//商品新增方法
	//Post请求
	//实现新增商品
}

func (c *GoodsController) UpdateGet() {
	//商品修改方法
	//Get请求
	//渲染 修改商品页面
}

func (c *GoodsController) UpdatePost() {
	//商品修改方法
	//Post请求
	//实现修改商品
}

func (c *GoodsController) DetailGet() {
	//商品单个查询方法
	//Get请求
	//渲染 商品查询页面
}

func (c *GoodsController) DetailPost() {

	//商品单个查询方法
	//Post请求
	//实现查询商品
}
