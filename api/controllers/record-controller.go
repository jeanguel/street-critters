package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jeanguel/street-critters/api/config"
	"github.com/jeanguel/street-critters/api/models"
	"github.com/jeanguel/street-critters/api/repositories"
	"github.com/jeanguel/street-critters/api/utils"
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

func GetRecordsByBoundingBox(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	vars := mux.Vars(r)
	startLongitude, err := strconv.ParseFloat(vars["startLong"], 64)
	if err != nil {
		config.MainLogger.Warn.Println("cannot parse start longitude variable:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Invalid start longitude provided"}
	}
	startLatitude, err := strconv.ParseFloat(vars["startLat"], 64)
	if err != nil {
		config.MainLogger.Warn.Println("cannot parse start latitude variable:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Invalid start latitude provided"}
	}
	endLongitude, err := strconv.ParseFloat(vars["endLong"], 64)
	if err != nil {
		config.MainLogger.Warn.Println("cannot parse end longitude variable:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Invalid end longitude provided"}
	}
	endLatitude, err := strconv.ParseFloat(vars["endLat"], 64)
	if err != nil {
		config.MainLogger.Warn.Println("cannot parse end latitude variable:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Invalid end latitude provided"}
	}

	records, err := repositories.GetRecordsByBoundingBox(
		startLongitude,
		startLatitude,
		endLongitude,
		endLatitude,
	)
	if err != nil {
		config.MainLogger.Warn.Println("repository fetch error:", err.Error())
		return 400, models.BaseAPIResponse{Message: "Unable to retrieve records"}
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

	place, err := utils.GetPlaceFromGeocode(recordSchema.Geopoint[0], recordSchema.Geopoint[1])
	if err != nil {
		config.MainLogger.Warn.Println(
			"unable to retrieve place details, resorting to placeholder:",
			err.Error(),
		)
		recordSchema.Address = "Placeholder"
	} else {
		recordSchema.Address = place
	}

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
