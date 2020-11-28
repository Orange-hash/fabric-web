package models

type Goods struct {
	GId     int    `orm:"pk;auto"`
	GName   string `orm:"size(100)"`
	GSize   string `orm:"size(100)"`
	GInfo   string
	GImgUrl string `orm:"size(100)"`
}

type Admin struct {
	AId       int    `orm:"pk;auto"`
	AName     string `orm:"size(64)"`
	AAccount  string `orm:"size(16)"`
	APassword string `orm:"size(32)"`
	AImageUrl string `orm:"size(255)"`
}
