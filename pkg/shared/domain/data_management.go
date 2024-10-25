package domain

import (
	"time"
)

// DataManagement model
type DataManagement struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Type      string    `gorm:"column:type;type:varchar(255)" json:"type"`
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Position  string    `gorm:"column:position;type:varchar(255)" json:"position"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of DataManagement model
func (DataManagement) TableName() string {
	return "data_managements"
}
