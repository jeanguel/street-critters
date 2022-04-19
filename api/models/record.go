package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/mitchellh/mapstructure"

	"github.com/jeanguel/street-critters/api/config"
)

type Record struct {
	ID          int64        `mapstructure:"record_id"`
	ReferenceID string       `mapstructure:"record_reference_id"`
	Type        RecordType   `mapstructure:"record_type"`
	Notes       *string      `mapstructure:"record_notes"`
	Geopoint    pgtype.Point `mapstructure:"record_geo_point"`
	Address     string       `mapstructure:"record_address"`
	CreatedAt   time.Time    `mapstructure:"created_at"`
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
