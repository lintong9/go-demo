package main

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/malisit/kolpa"
	uuid "github.com/satori/go.uuid"
)

func init() {
	_ = initDB()
}
func main() {
	// ctx := context.WithValue(context.Background(),"test","test")
	// Redis.String(ctx)
	k := kolpa.C()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 20; i++ {
		go func() {
			insertDb(k, &wg)
		}()
	}
	wg.Wait()
	fmt.Println("finish")
	return
}

var DB *sqlx.DB
var times int

func initDB() (err error) {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/golang?charset=utf8"
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("connecting to MySQL...")
	return
}

func insertDb(k kolpa.Generator, wg *sync.WaitGroup) {
	// stmt, _ := DB.Prepare(sqlStr)
	for {
		sqlStr := "INSERT INTO go_uuid (uuid, created_time) VALUES"
		for i := 0; i < 500; i++ {
			str, _ := uuid.NewV4()
			insertTime := time.Now().Unix()
			if i == 0 {
				sqlStr += fmt.Sprintf("('%s', %d)", str, insertTime)
			} else {
				sqlStr += fmt.Sprintf(",('%s', %d)", str, insertTime)
			}
		}
		_, err := DB.Exec(sqlStr)
		if err != nil {
			fmt.Printf("exec failed, err:%v\n", err)
			return
		}
		fmt.Println("插入完成")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Microsecond)
	}
}

func getRedis() {

}
