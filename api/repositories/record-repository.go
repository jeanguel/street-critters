package repositories

import (
	"fmt"

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

	for i, result := range rawResult {
		record := models.Record{}
		if err := mapstructure.Decode(result, &record); err != nil {
			return results, fmt.Errorf("failure on item %d: %s", i, err.Error())
		}

		results = append(results, record)
	}

	return results, nil
}
