package models

import (
	. "fiber/models/mysql"
	"time"
)

// @docs https://gorm.io/docs/connecting_to_the_database.html
// su dung json:"omitempty" để ẩn những giá trị bị thiếu khi trả về json
type RdUser struct {
	Uid uint32 `gorm:"primary_key"`
	Gid uint32 // int32 ~ type int(11) in mysql
	Project uint32  // int32 ~ type int(11) in mysql
	Fullname string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	InviteDate time.Time
	VietID uint32 `gorm:"column:VietID"`
	UserProjects []UserProjects `gorm:"foreignKey:Uid;"`
	UserField []UserField `gorm:"foreignKey:Uid;"`
}

// Set User's table name to be `rd_user`
func (RdUser) TableName() string {
	return "rd_user"
}

func GetUserByUid(uid uint32) (result []RdUser, count int) {
	DB.Debug().Select("*").Where(&RdUser{Uid: uid}).Take(&result)
	return result, len(result)
}

func GetUserByUidMongo(uid uint32) {

}