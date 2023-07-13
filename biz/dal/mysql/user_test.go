package mysql

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"testing"
)

func TestCreateUser(t *testing.T) {
	_ = Init()
	u := &User{
		UserName: "test",
		Password: "123456",
	}
	user_id, err := CreateUser(u)
	if err != nil {
		fmt.Printf("%v", false)
		return
	}
	fmt.Printf("%v", user_id)
}

func TestQueryUserById(t *testing.T) {
	_ = Init()
	user, err := QueryUserById(1)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("%v", user)
}

func TestQueryUsers(t *testing.T) {
	_ = Init()
	users, err := QueryUsers()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("%v", users)
	bytes, _ := json.Marshal(users)
	fmt.Printf("%s\n", string(bytes))
}
