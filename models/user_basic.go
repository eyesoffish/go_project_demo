package models

import (
	"go_chat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	Identity      string
	ClientPort    string
	LoginTime     *time.Time
	HeartBeatTime *time.Time
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.Db.Find(&data)
	// for _, v := range data {
	// 	fmt.Println("查到的数据: ", v)
	// }
	return data
}
