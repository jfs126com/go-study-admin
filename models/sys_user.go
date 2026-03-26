package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(20);" json:"userName"`
	PassWord  string `gorm:"column:password;type:varchar(100);" json:"passWord"`
	Phone     string `gorm:"column:phone;type:varchar(11);" json:"phone"`
	WxUnionID string `gorm:"column:wx_union_id;type:varchar(255);" json:"wxUnionID"`
	WxOpenID  string `gorm:"column:wx_open_id;type:varchar(255);" json:"wxOpenID"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
}

func (table *SysUser) TableName() string {
	return "sys_user"
}

func GetUserByUsernamePass(username, pass string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("username = ? and password = ?", username, pass).Find(data).Error
	return data, err
}
