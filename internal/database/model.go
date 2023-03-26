package database

type User struct {
	UserID string `gorm:"column:user_id"`
	Email  string `gorm:"column:email"`
	Name   string `gorm:"column:name"`
}

func (_ *User) Tablename() string {
	return "tbl_user"
}
