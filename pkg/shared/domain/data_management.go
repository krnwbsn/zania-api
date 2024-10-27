package domain

import (
	"time"
)

// DataManagement model
type DataManagement struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Type      string    `gorm:"column:type;type:varchar(255)" json:"type"`
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Position  int       `gorm:"column:position;type:int" json:"position"`
	ImageUrl  string    `gorm:"column:image_url;type:varchar(255)" json:"imageUrl"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName return table name of DataManagement model
func (DataManagement) TableName() string {
	return "data_managements"
}
