package data

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewAppDB(t *testing.T) {
	if db := NewAppDB(); db == nil {
		t.Error("Open db error")
	}
}

func TestSQLStatement(t *testing.T) {
	database := NewAppDB()
	if database == nil {
		t.Error("Create new int_eco db error")
	}

	const TableName = "token_info"
	sql := database.generateCreateSQL(TableName)
	if sql != "CREATE TABLE `token_info`(`id` int(64) unsigned PRIMARY KEY AUTO_INCREMENT,`access_token` varchar(50),`expires_in` timestamp not null,`refresh_token` varchar(50),`refresh_token_expires_in` timestamp,`scope` VARCHAR(255),`owner_id` varchar(100),`endpoint_id` varchar(100));" {
		t.Errorf("Create statement error: %s\n", sql)
	}
}

func TestConn(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Connect to db error")
		}
	}()

	database := NewAppDB()
	database.Conn()
}
func TestCreateTable(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("Open db error %v\n", err)
		}
	}()

	database := NewAppDB()
	database.Conn()
	database.CreateTokenTable("token_info")

	if ret := database.TableExists("token_info"); ret != true {
		t.Error("table `token_info` does not exist")
	}
}
func TestDropTable(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Error")
		}
	}()

	database := NewAppDB()
	if database == nil {
		t.Errorf("create table failed")
	}

	const TableName = "token_info"
	if ret := database.TableExists(TableName); ret != true {
		t.Log("table `token_info` does not exist")
	}

	if err := database.DropTable(TableName); err != nil {
		t.Errorf("drop table: [%s] failed", TableName)
	}
}
