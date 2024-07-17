package svc

import (
	"CatDisk/core/internal/config"
	"CatDisk/core/models"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c),
		RDB:    models.InitRedis(c),
	}
}
