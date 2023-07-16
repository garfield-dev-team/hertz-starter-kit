package dal

import (
	"fmt"

	"hertz-starter-kit/biz/dal/mysql"
)

func SetupConn() {
	if err := mysql.Init(); err != nil {
		panic(fmt.Errorf("fail to connect database: %w", err))
	}

	if err := mysql.Migrate(); err != nil {
		panic(fmt.Errorf("fail to migrate schema: %w", err))
	}
}
