package models

type User struct {
	IDUser   uint   `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}