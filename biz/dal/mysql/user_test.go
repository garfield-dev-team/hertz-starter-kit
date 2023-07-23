package mysql

import (
	"encoding/json"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "Add user 1",
			args: args{user: &User{UserName: "测试用户2333", Password: "2333"}},
		},
		{
			name: "Add user 2",
			args: args{user: &User{UserName: "测试用户666", Password: "666"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("CreateUser() success with result: %v", got)
		})
	}
}

func TestQueryUserById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Get user 1",
			args: args{id: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryUserById() success with result\n %s", string(res))
		})
	}
}

func TestQueryUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    []*User
		wantErr bool
	}{
		{
			name: "Get all users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryUsers() success with result\n %s", string(res))
		})
	}
}
