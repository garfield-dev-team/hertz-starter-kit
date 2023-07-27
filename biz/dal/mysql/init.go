package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hertz-starter-kit/pkg/config"
)

// 正确处理 `time.Time` 需要带上 parseTime 参数
// 支持完整的 UTF-8 编码，需要用 charset=utf8mb4
// 还有一个 collation 参数无需设置，默认就是 utf8mb4_general_ci
// 注意，dbname 是 docker-compose 的 MYSQL_DATABASE
//var dsn = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() (err error) {
	// 本地调试的时候没有读取配置，`config.Config` 空指针解引用会 panic
	// 这里设置一个默认值，给本地调试用
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 设置 GORM 默认日志级别
	level := logger.Error
	if config.Config != nil {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Database.UserName,
			config.Config.Database.Password,
			config.Config.Database.Host,
			config.Config.Database.Port,
			config.Config.Database.Name,
		)
		// 启用 debug 模式，打印 SQL 语句
		if config.Config.Server.RunMode == "debug" {
			level = logger.Info
		}
	}
	// GORM 配置参考
	// https://gorm.io/zh_CN/docs/gorm_config.html
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）
		// 如果没有这方面的要求，可以在初始化时跳过默认事务，这将获得大约 30%+ 性能提升
		SkipDefaultTransaction: true,
		// 启用缓存以提高效率
		PrepareStmt: true,
		// 打印 GORM 为我们生成的 SQL
		Logger: logger.Default.LogMode(level),
	})
	return
}

func Migrate() error {
	// 可以一次性将多个模型结构体对应的表进行创建或更新，比如：
	// DB.AutoMigrate(&model1{}, &model2{}, &model3{})
	// 注意，当模型结构体中有外键关联时，需要按照依赖关系的顺序进行创建
	// 比如，model1 有一个外键关联到 model2，则需要先创建 model2 的表，再创建 model1 的表

	// 如果两张表有互相依赖，比如 many to many，则任意顺序创建即可
	// 注意，many to many 实际上依赖的是一张连接表
	return DB.AutoMigrate(&User{}, &Article{}, &Category{})
}
