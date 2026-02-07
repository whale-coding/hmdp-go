package mysql

import (
	"fmt"
	"hmdp-go/internal/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMySQL 初始化数据库连接，返回数据库实例
func InitMySQL() *gorm.DB {
	// 读取配置文件中的数据库连接
	cfg := config.AppConfig.MySQL

	// 构造 dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect MySQL: %v", err)
	}

	log.Println("Successfully connected to MySQL")

	// 拿到通用的数据库对象，做一些额外的数据库配置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间
	
	return db
}
