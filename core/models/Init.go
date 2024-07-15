package models

import (
	"log"
	"xorm.io/xorm"
)

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:0508@/cloud-disk?charset=utf8")

	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return engine
}
