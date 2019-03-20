package data

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewIntEcoDB(t *testing.T) {
	if db := NewIntEcoDB(); db == nil {
		t.Error("Open db error")
	}
}

func TestSQLStatement(t *testing.T) {
	database := NewIntEcoDB()
	if database == nil {
		t.Error("Create new int_eco db error")
	}

	const TableName = "token_info"
	sql := database.generateCreateSQL(TableName)
	if sql != "DROP TABLE IF EXISTS `token_info`; CREATE TABLE `token_info`(`id` int(64) unsigned PRIMARY KEY AUTO_INCREMENT,`access_token` varchar(50),`expires_in` timestamp not null,`refresh_token` varchar(50),`refresh_token_expires_in` timestamp,`scope` VARCHAR(255),`owner_id` varchar(100),`endpoint_id` varchar(100));" {
		t.Errorf("Create statement error: %s\n", sql)
	}
}

func TestConn(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Connect to db error")
		}
	}()

	database := NewIntEcoDB()
	database.Conn()

	// conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:7002)/inteco")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// err = conn.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }
}
func TestCreateTable(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("Open db error %v\n", err)
		}
	}()

	database := NewIntEcoDB()
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

	database := NewIntEcoDB()
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
