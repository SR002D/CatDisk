package test

import (
	"CatDisk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestAddUser(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:0508@/cloud-disk?charset=utf8")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)

	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _ := json.Marshal(data)
	fmt.Printf("%s", resp)
}
