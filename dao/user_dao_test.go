package dao

import (
    "testing"
    "github.com/didi/gendry/scanner"
    "fmt"
    "github.com/didi/gendry/manager"
    "time"
    "database/sql"
    "github.com/didi/gendry/builder"
    "context"
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

func TestUpdateUser(t *testing.T) {
    result, err := builder.AggregateQuery(context.Background(), conn, "user", nil, builder.AggregateCount("*"))
    t.Log(result, err, result.Int64(), result.Float64())

    //where := map[string]interface{}{
    //    "id": 1,
    //}
    //cond, vals, err := builder.BuildUpdate("people", where, map[string]interface{}{"age": "`age + 1`"})
    //t.Log(cond, vals, err)
    //res, err := conn.Exec(cond, vals...)
    //t.Log(res.RowsAffected())


    //res,err := conn.Exec("update people set age=age+? where id=?",[]interface{}{1,1})
    //t.Log(res.RowsAffected())

    //stm,err:=conn.Prepare("update people set age = age + ? where id = ?")
    //res,err:=stm.Exec(1,1)
    //t.Log(res.RowsAffected())

}

func TestCount(t *testing.T) {
    cond, vals, err := builder.BuildSelect("user", nil, []string{"count(*) as num"})
    t.Log(cond, vals, err)
    rows, err := conn.Query(cond, vals...)

    type result struct {
        Num int `ddb:"num"`
    }
    var rel []result
    for rows.Next(){
        var num int
        rows.Scan(&num)
        t.Log(rel,num)
    }
}
