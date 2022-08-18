package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	DBEngin *gorm.DB
	ReEngin *redis.Client
)
