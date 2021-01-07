package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"jwt/utils"
	"os"
	"time"

	"log"
)

var (
	DB *gorm.DB
	c  = *utils.NewCfg()
)

//初始化sql
func InitSQL() (DB *gorm.DB) {
	cfg := c.InitConfig()
	dsn := cfg.GetString("MySQL.DataSource")

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "db_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `we_user`
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,   // 慢 SQL 阈值
				LogLevel:      logger.Silent, // Log level
				Colorful:      false,         // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic("MySQL启动异常")
	}

	if err := DB.AutoMigrate(new(User)); err != nil {
		log.Println("同步数据库表失败:", err)

	}
	return DB
}
