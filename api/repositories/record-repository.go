package repositories

import (
	"github.com/jackc/pgtype"
	"github.com/mitchellh/mapstructure"

	"github.com/jeanguel/street-critters/api/config"
	"github.com/jeanguel/street-critters/api/models"
)

func GetAllRecords(searchQuery *models.SearchQuery) ([]models.Record, error) {
	results := []models.Record{}
	params := []interface{}{
		searchQuery.Amount,
		searchQuery.GetOffset(),
	}

	rawResult, err := config.ExecPSQLFunc("report.get_all_records", params...)
	if err != nil {
		return results, err
	}

	err = mapstructure.Decode(rawResult, &results)
	for i, result := range results {
		results[i].FloatGeopoint = [2]float64{result.Geopoint.P.X, result.Geopoint.P.Y}
	}

	return results, err
}

func GetRecordsByBoundingBox(startLongitude, startLatitude, endLongitude, endLatitude float64) ([]models.Record, error) {
	results := []models.Record{}

	boundingBox := pgtype.Box{
		P: [2]pgtype.Vec2{
			{X: startLongitude, Y: startLatitude},
			{X: endLongitude, Y: endLatitude},
		},
		Status: pgtype.Present,
	}

	rawResult, err := config.ExecPSQLFunc("report.get_records_by_bounding_box", boundingBox)
	if err != nil {
		return results, err
	}

	err = mapstructure.Decode(rawResult, &results)
	for i, result := range results {
		results[i].FloatGeopoint = [2]float64{result.Geopoint.P.X, result.Geopoint.P.Y}
	}

	return results, err
}
