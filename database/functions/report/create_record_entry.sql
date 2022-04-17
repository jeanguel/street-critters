CREATE OR REPLACE FUNCTION report.create_record_entry (
  p_record_reference_id VARCHAR,
  p_record_type report.record_type,
  p_record_notes TEXT,
  p_record_geo_point POINT,
  p_record_address TEXT
) RETURNS TABLE (
  created_record_id BIGINT,
  success BOOLEAN,
  message TEXT
) AS
$func$
DECLARE
  v_created_record_id BIGINT := 0;
  v_success BOOLEAN := TRUE;
  v_message TEXT := '';
BEGIN
  INSERT INTO report.record (
    record_reference_id,
    record_type,
    record_notes,
    record_geo_point,
    record_address
  ) VALUES (
    p_record_reference_id,
    p_record_type,
    p_record_notes,
    p_record_geo_point,
    p_record_address
  );

  RETURN QUERY SELECT v_created_record_id, v_success, v_message;
END;
$func$ LANGUAGE plpgsql;
