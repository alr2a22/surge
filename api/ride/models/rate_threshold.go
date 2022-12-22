package models

import (
	"surge/internal/db"
)

type RideRateThreshold struct {
	db.Model
	Threshold   *uint64  `json:"threshold" gorm:"unique" validate:"required,number,gte=0"`
	Coefficient *float32 `json:"coefficient" validate:"required,numeric,gte=1"`
}

type RideRateThresholdList []RideRateThreshold

func (m *RideRateThresholdList) List() error {
	DB := db.GetDBConn()
	return DB.Order("threshold").Find(m).Error
}

func (m *RideRateThreshold) GetByID(id int) error {
	DB := db.GetDBConn()
	return DB.First(m, id).Error
}

func (m *RideRateThreshold) Create() error {
	DB := db.GetDBConn()
	return DB.Create(m).Error
}

func (m *RideRateThreshold) Save() error {
	DB := db.GetDBConn()
	return DB.Save(m).Error
}

func (m *RideRateThreshold) DeleteByID(id int) error {
	DB := db.GetDBConn()
	return DB.Delete(m, id).Error
}

func GetCurrentCoefficient(nReq int64) (float32, error) {
	DB := db.GetDBConn()
	var rrt RideRateThreshold
	err := DB.Order("threshold DESC").Take(&rrt, "threshold <= ?", nReq).Error
	if err != nil {
		return 0, err
	}
	return *rrt.Coefficient, nil
}
