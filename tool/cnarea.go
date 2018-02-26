package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "../faker/data/cnarea.db"

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
		area_code VARCHAR(64),
		zip_code VARCHAR(64),
		city_code VARCHAR(64),
		areaname VARCHAR(64),
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
	db, err := sql.Open("mysql", "evi1:w5nbb@tcp(192.168.1.108:3306)/fakerhub?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 统计条数
	countSql := "SELECT count(1) num FROM cnarea_2016 WHERE level = 4"
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
	querySql := "SELECT area_code, zip_code, city_code, name, short_name, merger_name, lng, lat FROM cnarea_2016 WHERE level = 4"
	rows, err := db.Query(querySql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var area_code, zip_code, city_code, name, short_name, merger_name, lng, lat string
		if err := rows.Scan(&area_code, &zip_code, &city_code, &name, &short_name, &merger_name, &lng, &lat); err != nil {
			log.Fatal(err)
		}
		tempMergerName := strings.Split(merger_name, ",")
		if len(tempMergerName) == 0 {
			tempMergerName[0] = "地市地址拆分错误"
		} else {
			tempMergerName = tempMergerName[0 : len(tempMergerName)-1]
		}
		areaname := strings.Join(tempMergerName, ",")
		//		fmt.Println(area_code, zip_code, city_code, areaname, name, short_name, lng, lat)
		row := make(map[string]string)
		row["area_code"] = area_code
		row["zip_code"] = zip_code
		row["city_code"] = city_code
		row["areaname"] = areaname
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
	stmt, err := tx.Prepare("insert into cnarea_2016(area_code, zip_code, city_code, areaname, name, short_name, lng, lat) values(?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	for _, row := range data {
		_, err = stmt.Exec(row["area_code"], row["zip_code"], row["city_code"], row["areaname"], row["name"], row["short_name"], row["lng"], row["lat"])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func main() {
	initSqlite()
	mysqlData, dataNum := exetractMySQL()
	fmt.Printf("total rows: %d\n", dataNum)
	loadSqlite(mysqlData)
}
