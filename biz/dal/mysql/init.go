package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 正确处理 `time.Time` 需要带上 parseTime 参数
// 支持完整的 UTF-8 编码，需要用 charset=utf8mb4
// 注意，dbname 是 docker-compose 的 MYSQL_DATABASE
var dsn = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,                                // 常规 CRUD 操作跳过默认事务
		PrepareStmt:            true,                                // 启用缓存以提高效率
		Logger:                 logger.Default.LogMode(logger.Info), // 打印 GORM 为我们生成的 SQL
	})
	return
}
