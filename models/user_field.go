package models

// @docs https://gorm.io/docs/connecting_to_the_database.html
type UserField struct {
	Ufid uint32 `gorm:"primary_key"`
	Uid uint32
	Ukey string
	Uvalue string
}
