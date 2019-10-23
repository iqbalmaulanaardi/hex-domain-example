package models

type User struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	CreatedAt int64  `gorm:"not null;" json:"createdAt"`
	UpdatedAt int64  `gorm:"not null;" json:"updatedAt"`
}
