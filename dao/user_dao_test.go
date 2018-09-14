package dao

import (
	"testing"
	"github.com/didi/gendry/scanner"
	"fmt"
	"github.com/didi/gendry/manager"
	"time"
	"database/sql"
)

var conn *sql.DB
var err error

func init() {
	conn, err = manager.New("test", "root", "123456", "127.0.0.1").Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(3306).Open(true)
}

func TestGetUsers(t *testing.T) {

	if err != nil {
		fmt.Println(err)
	}
	rows, _ := conn.Query("select * from user")
	defer rows.Close()
	result, _ := scanner.ScanMap(rows)
	for idx, record := range result {
		fmt.Println(idx, string(record["id"].([]byte)), string(record["name"].([]byte)), string(record["password"].
		([]byte)), string(record["email"].([]byte)))
	}

}
