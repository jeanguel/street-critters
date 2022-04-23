package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type reverseGeocodeResponse struct {
	ID          int64             `json:"place_id"`
	License     string            `json:"licence"`
	OSMType     string            `json:"osm_type"`
	OSMID       int64             `json:"osm_id"`
	DisplayName string            `json:"display_name"`
	Address     map[string]string `json:"address"`
	BoundingBox []string          `json:"boundingbox"`
}

// GetPlaceFromGeocode
func GetPlaceFromGeocode(longitude, latitude float64) (string, error) {
	params := url.Values{}
	params.Add("lon", fmt.Sprint(longitude))
	params.Add("lat", fmt.Sprint(latitude))
	params.Add("format", "json")

	response, err := http.Get("https://nominatim.openstreetmap.org/reverse?" + params.Encode())
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	place := reverseGeocodeResponse{}
	if err := json.Unmarshal(body, &place); err != nil {
		return "", err
	}

	return place.DisplayName, nil
}
