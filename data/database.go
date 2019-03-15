// DB service

package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB system db service
type IntEcoDB struct {
	db *sql.DB
}

// func (db *IntEcoDB) storeToken(tokenInfo *models.AuthInfo) error {

// }

// func (db *IntEcoDB) getToken() (info models.AuthInfo, err error) {

// }

// NewIntEcoDB create new db instance
func NewIntEcoDB() (dbInstance *IntEcoDB, err error) {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/inteco")
	if err != nil {
		log.Fatal(err)
	}

	// stmt, err := db.Prepare("")
	return &IntEcoDB{
		db: db,
	}, nil
}
