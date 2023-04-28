package initialization

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nft_platform/global"
	"time"
)

func initGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("初始化mysql出错了,err:%s", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("初始化mysql出错了,err:%s", err))
	}
	sqlDB.SetMaxIdleConns(global.Conf.Mysql.MaxIdle)
	sqlDB.SetMaxOpenConns(global.Conf.Mysql.MaxOpenIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Conf.Mysql.MaxLifeTime) * time.Second)
	global.DB = db
}
