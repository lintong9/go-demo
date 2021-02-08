package main

import (
	"fmt"
	"sync"
	"time"

	"demo/Excel"
	"demo/Snow"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/malisit/kolpa"
	// "demo/Excel"
)

func init() {
	_ = initDB()
}
func main() {
	Excel.GetTest()
	// ctx := context.WithValue(context.Background(),"test","test")
	// Redis.String(ctx)
	// k := kolpa.C()
	// wg := sync.WaitGroup{}
	// wg.Add(10)
	// for i := 0; i < 20; i++ {
	// 	go func() {
	// 		insertDb(k, &wg)
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("finish")
	return
	// getSnowFlake()
	// return
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
			str, _ := Snow.Snow.GetSnowflakeId()
			insertTime := time.Now().Unix()
			if i == 0 {
				sqlStr += fmt.Sprintf("('%d', %d)", str, insertTime)
			} else {
				sqlStr += fmt.Sprintf(",('%d', %d)", str, insertTime)
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

func getSnowFlake() {
	var count int = 1000000
	mapId := make(map[int64]int64, count)
	fmt.Println("start,count:", count)
	for i := 0; i < count; i++ {
		id, ts := Snow.Snow.GetSnowflakeId()
		mapId[id] = ts
	}
	fmt.Println("done,count:", count, ",mapCount:", len(mapId))
}
