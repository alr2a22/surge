package models

import (
	"surge/api/account/models"
	"surge/internal/db"
	"time"
)

type RideRequest struct {
	ID        uint `gorm:"primarykey"`
	Lat       float32
	Long      float32
	District  string
	CreatedAt time.Time `gorm:"index"`
	UserID    uint
	User      models.User
}

func (m *RideRequest) Create() error {
	DB := db.GetDBConn()
	return DB.Create(m).Error
}

func (m *RideRequest) CountRideRequestWithDistrictWithWindow(district string, w time.Duration) (int64, error) {
	DB := db.GetDBConn()
	var count int64
	err := DB.Where("district=? and created_at BETWEEN ? AND ?", district, time.Now(), time.Now().Add(w)).Count(&count).Error
	return count, err
}
