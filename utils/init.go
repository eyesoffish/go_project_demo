package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper err:", err)
	}
	fmt.Println("config app :", viper.Get("mysql"))
}

func InitMySql() {
	// 自定义日志模版, 打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢sql 的fa'zhi
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	var err error
	Db, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("db open fail")
	}
}
