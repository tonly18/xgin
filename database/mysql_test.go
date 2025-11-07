package database

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	dBConf := []*MySqlConfig{
		{
			Role:      "default",
			Host:      "54.69.237.139",
			Port:      8082,
			Dbname:    "test-sports-db",
			Username:  "root",
			Password:  "123456",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
	if err := initMysql(dBConf); err != nil {
		panic(err)
	}
}

type Test struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	itime string `json:"itime"`
}

func TestRedis(t *testing.T) {
	db := GetDB(RoleDefault)
	fmt.Printf("------db:%+v\n", db)
	//db.Clauses(dbresolver.Write)
	//db.Clauses(dbresolver.Read)

	mdb, err := db.DB()
	fmt.Printf("------err:%+v\n", err)
	fmt.Printf("------mdb:%+v\n", mdb)
	fmt.Printf("------mdb.Ping:%+v\n", mdb.Ping())

	var data Test
	if err := db.First(&data, "id=?", 2).Error; err != nil {
		log.Fatal(err)

	}

	fmt.Printf("------data:%+v\n", data)
}
