package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/mitchellh/mapstructure"

	"github.com/jeanguel/street-critters/api/config"
)

type Record struct {
	ID            int64        `mapstructure:"record_id" json:"id"`
	ReferenceID   string       `mapstructure:"record_reference_id" json:"referenceId"`
	Type          RecordType   `mapstructure:"record_type" json:"type"`
	Notes         *string      `mapstructure:"record_notes" json:"notes,omitempty"`
	Geopoint      pgtype.Point `mapstructure:"record_geo_point" json:"-"`
	FloatGeopoint [2]float64   `mapstructure:"-" json:"geopoint"`
	Address       string       `mapstructure:"record_address" json:"address"`
	CreatedAt     time.Time    `mapstructure:"created_at" json:"createdAt"`
}

// Save
func (r *Record) Save() error {
	params := []interface{}{
		r.ReferenceID,
		r.Type,
		r.Notes,
		r.Geopoint,
		r.Address,
	}

	result, err := config.ExecPSQLFunc("report.create_record_entry", params...)
	if err != nil {
		return err
	}
	if len(result) < 1 {
		return fmt.Errorf("no response from database")
	}

	response := struct {
		BaseDBResponse `mapstructure:",squash"`
		RecordID       int64 `mapstructure:"created_record_id"`
	}{}
	if err := mapstructure.Decode(result[0], &response); err != nil {
		return err
	}
	if !response.Success {
		return fmt.Errorf("%s: %s", "unable to save entry", response.Message)
	}

	r.ID = response.RecordID
	return nil
}
