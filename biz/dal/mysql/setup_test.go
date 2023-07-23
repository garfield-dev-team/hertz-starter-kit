package mysql

import (
	"fmt"
	"os"
	"testing"

	"hertz-starter-kit/pkg/config"
)

func TestMain(m *testing.M) {
	if err := os.Setenv("GO_ENV", "test"); err != nil {
		panic(err)
	}
	// 在这里执行初始化操作，例如建立数据库连接、加载测试数据等
	fmt.Println("===[1/3]配置文件初始化===")
	config.SetupConfig()
	fmt.Println("===[2/3]数据库链接初始化===")
	_ = Init()
	fmt.Println("===[3/3]开始运行测试用例===")
	// 运行测试
	code := m.Run()
	os.Exit(code)
}
