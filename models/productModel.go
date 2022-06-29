package models

type Product struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null;uniqueIndex" json:"title" form:"title" valid:"required~Title is Required"`
	Price      int       `gorm:"not null" json:"price" form:"price" valid:"required~Price is Required"`
	Stock      int       `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock is Required"`
	CategoryID int       
	Category *Category
	TimeModel
}
