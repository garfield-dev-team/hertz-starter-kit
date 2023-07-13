package mysql

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func CreateUser(user *User) (uint, error) {
	if err := DB.Create(user).Error; err != nil {
		return 0, err
	}
	// 返回插入数据的主键
	return user.ID, nil
}

func QueryUserById(user_id uint) (*User, error) {
	var user *User
	// GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象
	// 查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误

	// 为了避免 ErrRecordNotFound 错误，可以使用 Find
	// 比如 db.Limit(1).Find(&user)，Find 方法可以接受 struct 和 slice 的数据

	// 需要注意，对单个对象使用 Find 而不带 limit
	// db.Find(&user)将会查询整个表并且只返回第一个对象，这是性能不高并且不确定的。
	result := DB.Where("id = ?", user_id).Limit(1).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUsers() ([]*User, error) {
	var users []*User
	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
