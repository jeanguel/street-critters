CREATE OR REPLACE FUNCTION report.get_all_records(
  p_limit INT,
  p_offset INT
) RETURNS TABLE (
  record_id BIGINT,
  record_reference_id VARCHAR,
  record_type report.record_type,
  record_notes TEXT,
  record_geo_point POINT,
  record_address TEXT,
  created_at TIMESTAMPTZ
) AS
$func$
BEGIN
  RETURN QUERY
  SELECT
    r.record_id,
    r.record_reference_id,
    r.record_type,
    r.record_notes,
    r.record_geo_point,
    r.record_address,
    r.created_at
  FROM report.record r
  ORDER BY r.created_at DESC
  LIMIT p_limit
  OFFSET p_offset;
END;
$func$ LANGUAGE plpgsql;
