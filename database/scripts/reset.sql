BEGIN;

CREATE SCHEMA IF NOT EXISTS entity;
CREATE SCHEMA IF NOT EXISTS report;

-- Use dynamic SQL to drop all tables
DO
$dyna_sql$
DECLARE
  q record;
BEGIN
  FOR q IN
    SELECT
      format('DROP TABLE IF EXISTS %s.%s CASCADE;', t.table_schema, t.table_name) AS stmt
    FROM information_schema.tables t
    WHERE t.table_schema IN ('entity', 'report')
      AND t.table_type = 'BASE TABLE'
  LOOP
    EXECUTE q.stmt;
  END LOOP;
END;
$dyna_sql$ LANGUAGE plpgsql;

-- Use dynamic SQL to drop all functions
DO
$dyna_sql$
DECLARE
  q record;
BEGIN
  FOR q IN
    SELECT
      format('DROP FUNCTION IF EXISTS %s.%s (' || oidvectortypes(proargtypes) || ');', ns.nspname, pp.proname) AS stmt
    FROM pg_proc pp
    INNER JOIN pg_namespace ns
      ON pp.pronamespace = ns.oid
    WHERE ns.nspname IN ('entity', 'report')
  LOOP
    EXECUTE q.stmt;
  END LOOP;
END;
$dyna_sql$ LANGUAGE plpgsql;

-- User dynamic SQL to drop all custom types
DO
$dyna_sql$
DECLARE
  q record;
BEGIN
  FOR q IN
    SELECT
      format('DROP TYPE IF EXISTS %s.%s', ns.nspname, t.typname) AS stmt
    FROM pg_type t
    LEFT JOIN pg_catalog.pg_namespace ns
      ON ns.oid = t.typnamespace
    WHERE (t.typrelid = 0
      OR (
        SELECT
          c.relkind = 'c'
        FROM pg_catalog.pg_class c
        WHERE c.oid = t.typrelid
      ))
      AND NOT EXISTS(
        SELECT 1 FROM pg_catalog.pg_type el
        WHERE el.oid = t.typelem
          AND el.typarray = t.oid
      )
      AND ns.nspname NOT IN ('pg_catalog', 'information_schema')
  LOOP
    EXECUTE q.stmt;
  END LOOP;
END;
$dyna_sql$ LANGUAGE plpgsql;

COMMIT;
