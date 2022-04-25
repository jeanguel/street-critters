package models

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jackc/pgtype"
	"gopkg.in/dealancer/validate.v2"
)

type RecordSchema struct {
	Type     RecordType `json:"recordType"`
	Notes    *string    `json:"notes"`
	Geopoint [2]float64 `json:"geopoint" validate:"empty=false & eq=2"`
	Address  string     `json:"address"`
}

// Map
func (rs RecordSchema) Map(body io.ReadCloser) (*RecordSchema, error) {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&rs); err != nil {
		return &rs, err
	}

	err := validate.Validate(&rs)
	return &rs, err
}

// Validate
func (rs RecordSchema) Validate() error {
	if !rs.Type.IsValid() {
		return fmt.Errorf("unknown record type provided [%s]", rs.Type)
	}

	return nil
}

// LoadModel
func (rs RecordSchema) ToModel() (*Record, error) {
	// TODO: Reference ID byte length must be configurable
	referenceBytes := make([]byte, 10)
	_, err := rand.Read(referenceBytes)
	if err != nil {
		return nil, err
	}

	point := pgtype.Vec2{X: rs.Geopoint[0], Y: rs.Geopoint[1]}
	return &Record{
		// TODO: Reference ID prefix must be configurable
		ReferenceID: fmt.Sprintf("SC-R-%s", base32.StdEncoding.EncodeToString(referenceBytes)),
		Type:        rs.Type,
		Notes:       rs.Notes,
		Geopoint:    pgtype.Point{P: point, Status: pgtype.Present},
		Address:     rs.Address,
	}, nil
}
