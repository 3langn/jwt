package models


type Book struct {
	Id uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	UserID uint64 `gorm:"not null" json:"-"` // Field is ignored by this package.
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}