package dao

import (
	"altair-backend/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func newMysql() *gorm.DB {
	DatabaseUri := config.ServerConfig.DbUri

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true, // 使用单数表名
		},
		PrepareStmt:            true, // 预编译
		SkipDefaultTransaction: true, // 禁用默认事务
	}

	// 测试环境，输出sql日志
	if config.ServerConfig.RunMode == "debug" {
		configs.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,
			},
		)
	}

	db, err := gorm.Open(mysql.Open(DatabaseUri), configs)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return db

}
