package geo_query

import (
	"fmt"
	"surge/internal/db"
)

func FindDistrict(lat, long float32) (string, error) {
	DB := db.GetDBConn()
	var districtID string
	p := fmt.Sprintf("POINT(%v %v)", long, lat)
	err := DB.Raw(
		`SELECT id FROM districts WHERE ST_Contains(wkb_geometry, ST_Transform(ST_GeomFromText(?,4326), 4326));`,
		p,
	).Scan(&districtID).Error

	if err != nil {
		return "", err
	}

	return districtID, nil
}
