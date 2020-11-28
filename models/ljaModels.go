package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "good:Good123456@tcp(rm-bp1z81221k151q604qo.mysql.rds.aliyuncs.com:3306)/good")
	orm.RegisterModel(new(Goods), new(Admin), new(GoodsBatch), new(BatchLink))
	orm.RunSyncdb("default", false, true)
}
