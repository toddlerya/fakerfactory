package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "../faker/data/data.db"

func initSqlite() {
	// 连接数据库
	os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 建表
	createSql := `CREATE TABLE IF NOT EXISTS cnarea_2016 (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		level VARCHAR(64),
		area_code VARCHAR(64),
		zip_code VARCHAR(64),
		city_code VARCHAR(64),
		area_name VARCHAR(64),
		name VARCHAR(64),
		short_name VARCHAR(64),
		lng VARCHAR(64),
		lat VARCHAR(64)
	)`

	_, err = db.Exec(createSql)
	if err != nil {
		log.Panicf("%q: %s\n", err, createSql)
		return
	}
}

func exetractMySQL() ([]map[string]string, int) {
	db, err := sql.Open("mysql", "root:root@tcp(172.16.5.193:3306)/cnarea?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 统计条数
	countSql := "SELECT count(1) num FROM cnarea_2016"
	var num string
	row := db.QueryRow(countSql)
	err = row.Scan(&num)
	if err != nil {
		log.Fatal(err)
	}
	dataNumber, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("共计%d条数据\n", dataNumber)

	var data []map[string]string

	// 查询
	//	querySql := "SELECT area_code, zip_code, city_code, name, short_name, merger_name, lng, lat FROM cnarea_2016 WHERE level = 4"
	querySql := "SELECT level, area_code, zip_code, city_code, name, short_name, merger_name, lng, lat FROM cnarea_2016"
	rows, err := db.Query(querySql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var level, area_code, zip_code, city_code, name, short_name, merger_name, lng, lat string
		if err := rows.Scan(&level, &area_code, &zip_code, &city_code, &name, &short_name, &merger_name, &lng, &lat); err != nil {
			log.Fatal(err)
		}
		area_name := merger_name
		//		fmt.Println(level, area_code, zip_code, city_code, area_name, name, short_name, lng, lat)
		row := make(map[string]string)
		row["level"] = level
		row["area_code"] = area_code
		row["zip_code"] = zip_code
		row["city_code"] = city_code
		row["area_name"] = area_name
		row["name"] = name
		row["short_name"] = short_name
		row["lng"] = lng
		row["lat"] = lat
		data = append(data, row)
	}
	return data, dataNumber
}

func loadSqlite(data []map[string]string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into cnarea_2016(level, area_code, zip_code, city_code, area_name, name, short_name, lng, lat) values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	for _, row := range data {
		_, err = stmt.Exec(row["level"], row["area_code"], row["zip_code"], row["city_code"], row["area_name"], row["name"], row["short_name"], row["lng"], row["lat"])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func main() {
	startTime := time.Now()
	initSqlite()
	mysqlData, dataNum := exetractMySQL()
	fmt.Printf("total rows: %d\n", dataNum)
	loadSqlite(mysqlData)
	costTime := time.Now().Sub(startTime)
	fmt.Println("MySQL-->SQLite 迁移数据耗时", costTime)
}
