package models

import (
	"DellDaven_API/repository"
	"github.com/astaxie/beego/orm"
)

//FrontText
type FrontText struct {
	Id       int    `json:"id"`
	Language string `orm:"column(language)" json:"language"`
	Key      string `orm:"column(key)" json:"key"`
	Value    string `orm:"column(value)" json:"value"`
}

func init() {
	orm.RegisterModel(new(FrontText))
}

func GetText(key string) (frontText FrontText, err error) {
	db := repository.GetSession()
	err = db.QueryTable(frontText).Filter("key", key).One(&frontText)
	return frontText, err
}
