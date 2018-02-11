# ginapp
gin框架的应用

## 与mysql的交互应用

### 安装mysql通用库

```bash
go get github.com/go-sql-driver/mysql
```

### 引入包连接数据库操作

```go
db.gopackage main

import (
	//依赖包https://github.com/go-sql-driver/mysql
	_ "../mysql"
	"database/sql"
	"fmt"
)

func main() {
	//打开连接
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	checkError(err)

	//插入数据
	stmt, err := db.Prepare("insert user set name=?,password=?,email=?")
	checkError(err)

	res, err := stmt.Exec("1111", "1111", "1111")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update user set name=? where id=?")
	checkError(err)

	res, err = stmt.Exec("2222", id)
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("select * from user")
	checkError(err)

	//遍历查询结果
	for rows.Next() {
		var id int
		var name string
		var password string
		var email string
		err = rows.Scan(&id, &name, &password, &email)
		checkError(err)
		fmt.Println(id, name, password, email)
	}

	//删除数据
	stmt, err = db.Prepare("delete from user where id=?")
	checkError(err)

	res, err = stmt.Exec(id)
	affect, err = res.RowsAffected()
	checkError(err)

	fmt.Println(affect)
	db.Close()

	m := map[string]string{
		"name": "aa",
		"age":  "30",
	}
	fmt.Println(m)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
```

## 与memcached的交互应用

### 安装memecache客户端能用库

```bash
go get github.com/bradfitz/gomemcache/memcache
```

### 引入库连接服务器操作

```go
package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// 连接服务器
	mc := memcache.New("127.0.0.1:11211")

	// 写数据 &memcache.Item{Key,Value,Flags,Expiration,casid}
	mc.Set(&memcache.Item{Key: "key_one", Value: []byte("hello")})
	mc.Set(&memcache.Item{Key: "key_two", Value: []byte("bill")})

	// 获取数据
	val, err := mc.Get("key_one")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("-- %s", val.Value)

	// 获取多个key,返回map
	it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range it {
		fmt.Printf("## %s => %s\n", k, v.Value)
	}
}
```

## 与redis的交互应用

### 安装redis客户端库

```bash
go get -u github.com/go-redis/redis
```

### 引入库建立连接并操作

```go
package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	// 新建客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	
	defer client.Close()
	// 连接服务器测试
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// 添加key-value
	err = client.Set("key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取存在的key
	val, err := client.Get("key1").Result()
	fmt.Println(val, err)

	// 获取不存在的key
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
```

## 与mongod的交互应用

### 安装mongod客户端库

```bash
go get gopkg.in/mgo.v2
```

### 建立会话进行读写

```go
package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	// 建立一个会话
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	// 延迟最后关系会话
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// 指定数据库与表
	c := session.DB("test").C("people")
	// 播放数据
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	// 查找数据
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

```
