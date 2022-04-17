CREATE TYPE report.record_type AS ENUM (
  'poop',
  'dead_animal'
);

CREATE TABLE IF NOT EXISTS report.record(
  record_id BIGSERIAL NOT NULL,
  record_reference_id VARCHAR(31) NOT NULL,
  record_type report.record_type NOT NULL DEFAULT 'poop',
  record_notes TEXT NULL,
  record_geo_point POINT NOT NULL,
  record_address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT PK_RECORD_RECORD_ID PRIMARY KEY(record_id),
  CONSTRAINT UNIQ_RECORD_RECORD_REFERENCE_ID UNIQUE(record_reference_id)
);
