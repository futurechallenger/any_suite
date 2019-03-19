// DB service

package data

import (
	"database/sql"
	"errors"
	"fmt"
	"int_ecosys/models"

	// Register mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// IntEcoDB system db service
type IntEcoDB struct {
	db *sql.DB
}

// StoreToken stores auth info including auth token and refresh token, etc.
func (database *IntEcoDB) StoreToken(tokenInfo *models.AuthInfo) (rowID int64, err error) {
	if tokenInfo == nil {
		return -1, errors.New("tokenInfo is invalid")
	}

	const insertStatement = `insert into token_info (
			access_token, expires_in, refresh_token,
			refresh_token_expires_in, scope, owner_id, 
			endpoint_id
		) values (?,?,?,?,?,?,?)`
	stmt, err := database.db.Prepare(insertStatement)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	r, insertErr := stmt.Exec(tokenInfo.AccessToken, tokenInfo.ExpiresIn,
		tokenInfo.RefreshToken, tokenInfo.RefreshTokenExpiresIn,
		tokenInfo.Scope, tokenInfo.OwnerID, tokenInfo.EndPointID)

	if insertErr != nil {
		return -1, err
	}

	insertRowID, rowErr := r.LastInsertId()
	if rowErr != nil {
		return -1, fmt.Errorf("Can not get ID of inserted row, ERROR: %v", rowErr)
	}

	return insertRowID, nil
}

// TODO: get token by what? `owner_id` or anything else?
// func (database *IntEcoDB) getToken() (info models.AuthInfo, err error) {
//}

// NewIntEcoDB create new db instance
func NewIntEcoDB() (dbInstance *IntEcoDB) {
	db := &IntEcoDB{
		db: nil,
	}
	return db
}

// Close close inner db instance
func (database *IntEcoDB) Close() {
	if database.db != nil {
		database.db.Close()
	}
}

// Conn connect to mysql database
func (database *IntEcoDB) Conn() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:7002)/inteco")
	if err != nil {
		fmt.Printf("Connect to db error %v\n", err)
		panic(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("ping db error %v\n", err)
		panic(err.Error())
	}

	database.db = db
}

// DBExists check if db exists
func (database *IntEcoDB) dbExists(dbName string, intecoDB *IntEcoDB) bool {
	database.Conn()
	row := intecoDB.db.QueryRow(fmt.Sprintf("SELECT dbname FROM INFORMATION_SCHEMA.SCHEMATA WHERE SWHERE SCHEMA_NAME = '%s'", dbName))
	dbname := ""
	row.Scan(&dbname)

	defer database.Close()

	return dbname != ""
}

// IntEcoDBExists check if `inteco` db exists
func (database *IntEcoDB) IntEcoDBExists() bool {
	return database.dbExists("inteco", database)
}

// TableExists check if a table exists
func (database *IntEcoDB) TableExists(tableName string) bool {
	database.Conn()

	sql := fmt.Sprintf("SELECT 1 FROM %s LIMIT 1;", tableName)
	_, err := database.db.Query(sql)

	defer database.Close()

	return err == nil
}

// generateCreateSQL generate `token_info` state for now
// Later it should be a generate function to generate
// create table statement
func (database *IntEcoDB) generateCreateSQL(tableName string) string {
	pk := "`id` int(64) unsigned PRIMARY KEY AUTO_INCREMENT"
	accessToken := "`access_token varchar(50)"
	expiresIn := "`expires_in` timestamp not null"
	refreshToken := "`refresh_token` varchar(50)"
	refreshTokenExpiresIn := "`refresh_token_expires_in` timestamp"
	scope := "`scope` VARCHAR(255)"
	// tokenType := "`token_type` " `Bearer`
	ownerID := "`owner_id` varchar(100)"
	endPointID := "`endpoint_id` varchar(100)"

	createSQL := fmt.Sprintf("CREATE TABLE `%s`(%s,%s,%s,%s,%s,%s,%s,%s);",
		tableName, pk, accessToken,
		expiresIn, refreshToken,
		refreshTokenExpiresIn,
		scope, ownerID, endPointID)
	createState := fmt.Sprintf("DROP TABLE IF EXISTS `%s`; %s", tableName, createSQL)

	return createState
}

// CreateTokenTable if the db is first initialized
func (database *IntEcoDB) CreateTokenTable(tableName string) {
	database.Conn()

	createState := database.generateCreateSQL(tableName)
	stmt, err := database.db.Prepare(createState)

	if err != nil {
		// TODO: use log
		fmt.Printf("Create table [%s] failed %v\n", tableName, err)
		// return err
		panic(err)
	}
	// defer stmt.Close()
	defer database.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("Execute create statement [%s] failed %v\n", createState, err)
		// return err
		panic(err)
	}
}

// DropTable drop mysql table
func (database *IntEcoDB) DropTable(tableName string) error {
	database.Conn()
	dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	stmt, err := database.db.Prepare(dropSQL)
	if err != nil {
		fmt.Printf("Drop table: `%s` failed", tableName)
		return err
	}
	// defer stmt.Close()
	defer database.Close()

	_, err = stmt.Exec()

	return err
}
