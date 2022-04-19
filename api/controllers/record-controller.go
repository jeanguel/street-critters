package controllers

import (
	"net/http"

	"github.com/jeanguel/street-critters/api/config"
	"github.com/jeanguel/street-critters/api/models"
	"github.com/jeanguel/street-critters/api/repositories"
)

// GetRecords
func GetRecords(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	searchQuery, err := models.SearchQuery{}.ParseParams(r.URL.Query())
	if err != nil {
		config.MainLogger.Warn.Println("query param parsing error:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Cannot parse query params"}
	}

	records, err := repositories.GetAllRecords(searchQuery)
	if err != nil {
		config.MainLogger.Warn.Println("repository fetch error:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Unable to retrieve errors"}
	}

	return 200, records
}

// CreateNewRecord
func CreateNewRecord(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	recordSchema, err := models.RecordSchema{}.Map(r.Body)
	if err != nil {
		config.MainLogger.Warn.Println("schema loading error:", err.Error())
		return 400, models.BaseAPIResponse{Message: err.Error()}
	}

	// TODO: Must get the address via OpenStreetMaps geopoint
	recordSchema.Address = "Placeholder"

	record, err := recordSchema.ToModel()
	if err != nil {
		config.MainLogger.Warn.Println("model loading error:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Unable to create record"}
	}

	if err := record.Save(); err != nil {
		config.MainLogger.Warn.Println("record creation error:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Unable to create record"}
	}

	return 201, models.BaseAPIResponse{Message: "Record created"}
}
