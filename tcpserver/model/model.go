package model

import (
	"awesomeProject/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 用户
type User struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Profile    string `json:"profile"`
	CreatedOn  int32  `json:"created_on"`
	ModifiedOn int32  `json:"modified_on"`
}

func (u *User) TableName() string {
	return "user"
}

// 创建DBEngine
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.PassWord,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime))
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 获取用户信息
func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("username = ?", u.Username)
	err := db.Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 更新用户nickname
func (u User) UpdateNick(db *gorm.DB) error {
	db = db.Model(&User{}).Where("username = ?", u.Username)
	return db.Update(User{Nickname: u.Nickname, ModifiedOn: int32(time.Now().Unix())}).Error
}

// 更新用户头像
func (u User) UpdateProf(db *gorm.DB) error {
	db = db.Model(&User{}).Where("username = ?", u.Username)
	return db.Update(User{Profile: u.Profile, ModifiedOn: int32(time.Now().Unix())}).Error
}
