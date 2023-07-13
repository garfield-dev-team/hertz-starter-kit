package dal

import "hertz-starter-kit/biz/dal/mysql"

func SetupConn() error {
	return mysql.Init()
}
