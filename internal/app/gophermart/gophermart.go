package gophermart

import (
	"database/sql"
	store "gophermart/internal/app/store/sqlstore"
	"net/http"
)

const databaseDriverName = "postgres"

func Start(config *config) error {
	database, err := newDB(config.DatabaseURI)
	if err != nil {
		return err
	}

	defer database.Close()
	store := store.New(database)
	server := newServer(store, config.AccuralSystemAddress)

	return http.ListenAndServe(config.RunAddress, server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	database, err := sql.Open(databaseDriverName, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
