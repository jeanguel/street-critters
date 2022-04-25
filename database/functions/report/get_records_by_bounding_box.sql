CREATE OR REPLACE FUNCTION report.get_records_by_bounding_box(
  p_bounds BOX
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

  -- FIXME: This needs to be limited. One idea is that points that
  --  are really close relative to the size of the bounding box
  --  should be rendered as one, but not sure how to do that yet.
  WHERE p_bounds @> r.record_geo_point;
END;
$func$ LANGUAGE plpgsql;
