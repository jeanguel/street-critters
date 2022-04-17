package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/jackc/pgx/v4"
)

// SQLTLSCert holds TLS Certs for PSQL as the 'pgx'
// package does not include TLS Certs when reading
// the connection string
var SQLTLSCert *tls.Config

// ExecRawPSQLQuery executes a raw query that does not return
// a result. Will only return an error.
func ExecRawPSQLQuery(query string) error {
	conn, err := pgx.Connect(context.Background(), Config.Database.PSQLConnectionString)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())
	if len(Config.Database.TLSCert) > 0 {
		conn.Config().TLSConfig = SQLTLSCert
	}

	_, err = conn.Exec(context.Background(), query)
	return err
}

// ExecPSQLFunc is a convenience wrapper for executing database functions
func ExecPSQLFunc(funcName string, args ...interface{}) ([]*map[string]interface{}, error) {
	results := []*map[string]interface{}{}

	// NOTE: Extra space after parenthesis is meant to
	// handle no arguments since stmt will have its last
	// character removed before applying closing parenthesis
	stmt := "select * from " + funcName + "( "
	for i := 0; i < len(args); i++ {
		stmt += fmt.Sprintf("$%d,", i+1)
	}

	stmt = stmt[:len(stmt)-1] + ")"
	conn, err := pgx.Connect(context.Background(), Config.Database.PSQLConnectionString)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())
	if len(Config.Database.TLSCert) > 0 {
		conn.Config().TLSConfig = SQLTLSCert
	}

	rows, err := conn.Query(context.Background(), stmt, args...)
	if err != nil {
		return results, err
	}

	defer rows.Close()

	fields := rows.FieldDescriptions()
	for rows.Next() {
		mappedItem := map[string]interface{}{}

		vs, err := rows.Values()
		if err != nil {
			MainLogger.Warn.Println("Could not map results:", err)
			continue
		}

		for i, v := range vs {
			mappedItem[string(fields[i].Name)] = v
		}

		results = append(results, &mappedItem)
	}

	if err := rows.Err(); err != nil {
		return results, err
	}

	return results, nil
}

func initializeDBCerts() {
	if len(Config.Database.TLSCert) < 1 {
		return
	}

	caCert, err := ioutil.ReadFile(Config.Database.TLSCert)
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	SQLTLSCert = &tls.Config{
		RootCAs: caCertPool,
	}
}
