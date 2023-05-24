package dummydb

import (
	"database/sql"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
	"github.com/hangnadi/simple-api-project-2/lib/sqlt"
	"github.com/hangnadi/simple-api-project-2/lib/storage/database"
	"github.com/jmoiron/sqlx"
)

// New dummy database
func New(paramDb *sql.DB) database.Database {

	dbMock := sqlt.InitMocking(paramDb, 1)
	return &dummyDB{
		db: dbMock,
	}
}

func (f *dummyDB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return f.db.Queryx(query, args...)
}

func (f *dummyDB) Begin() (*sql.Tx, error) {
	return f.db.Begin()
}

func (f *dummyDB) Beginx() (*sqlx.Tx, error) {
	return f.db.Beginx()
}

func (f *dummyDB) Master() *sqlx.DB {
	return f.db.Master()
}

func (f *dummyDB) Get(dest interface{}, query string, args ...interface{}) error {
	return f.db.Get(dest, query, args...)
}

func (f *dummyDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return f.db.Exec(query, args...)
}
