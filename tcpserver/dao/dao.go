package dao

import (
	"awesomeProject/tcpserver/model"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	myEngine *gorm.DB
	reEngine *redis.Client
}

func New(myEngine *gorm.DB, reEngine *redis.Client) *Dao {
	return &Dao{myEngine: myEngine, reEngine: reEngine}
}

// 获取用户信息 Dao层
func (d *Dao) GetUser(username string) (model.User, error) {
	user := model.User{Username: username}
	return user.Get(d.myEngine)
}

// 更新nickname Dao层
func (d *Dao) UpdateNick(nickname string, username string) error {
	user := model.User{
		Username: username,
		Nickname: nickname,
	}
	return user.UpdateNick(d.myEngine)
}

// 更新头像 Dao层
func (d *Dao) UpdateProf(profile string, username string) error {
	user := model.User{
		Username: username,
		Profile:  profile,
	}
	return user.UpdateProf(d.myEngine)
}
