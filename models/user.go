package models

// https://gorm.io/docs/html
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"` // ignore this field when write and read with struct, exluded from JSON output
}
