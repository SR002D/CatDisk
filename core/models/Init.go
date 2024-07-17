package models

import (
	"core/internal/config"
	"log"
	"xorm.io/xorm"
)

func Init(c config.Config) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", c.Mysql.DataSource)

	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
