package controllers

import (
	"github.com/jeanguel/street-critters/api/config"
)

func ApplyRoutes() {
	recordBlueprint := config.Router.PathPrefix("/record").Subrouter()

	recordBlueprint.HandleFunc(
		"/",
		JSONRoute(GetRecords),
	).Methods("GET")

	recordBlueprint.HandleFunc(
		"/",
		JSONRoute(CreateNewRecord),
	).Methods("POST")

	recordBlueprint.HandleFunc(
		"/bounding-box/{startLong}/{startLat}/{endLong}/{endLat}",
		JSONRoute(GetRecordsByBoundingBox),
	).Methods("GET")
}
