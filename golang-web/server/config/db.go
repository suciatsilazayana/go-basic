package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	host := "localhost"

	port := "9990"

	user := "noobee"

	pass := "noobee"

	dbname := "bootcamp"

	dbs, err := getPostgres(host, port, user, pass, dbname)
	if err != nil {
		panic(err)
	}

	err = dbs.Ping()

	if err != nil {
		panic(err)
	}

	return dbs

}

func getPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	// membuat data source
	desc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbs, err := createConnection(desc)
	if err != nil {

		return nil, err
	}

	return dbs, nil
}

func createConnection(desc string) (*sql.DB, error) {
	// proses membuka koneksi
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
