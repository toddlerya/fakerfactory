package faker

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func AddressColumns(conn *sql.DB, columns ...string) []map[string]string {

	num := Number(1, 752233)
	cols := strings.Join(columns, ",")
	var err error
	var queryRows []map[string]string
	//	fmt.Println("columns: ", columns, "num: ", num, "sql", fmt.Sprintf("SELECT %s FROM cnarea_2016 WHERE uid = %d and level = 4", cols, num))
	if len(columns) != 0 {
		//		queryRows, err = QuerySqlite(conn, fmt.Sprintf("SELECT %s FROM cnarea_2016 WHERE uid = %d and level = 4", cols, num))
		queryRows, err = QuerySqlite(conn, fmt.Sprintf("SELECT %s FROM cnarea_2016 WHERE uid = %d", cols, num))
	} else {
		//		queryRows, err = QuerySqlite(conn, fmt.Sprintf("SELECT * FROM cnarea_2016 WHERE uid = %d and level = 4", num))
		queryRows, err = QuerySqlite(conn, fmt.Sprintf("SELECT * FROM cnarea_2016 WHERE uid = %d", num))
	}
	if err != nil {
		log.Fatal(err)
	}
	return queryRows
}

func Address(conn *sql.DB) map[string]string {
	colMap := AddressColumns(conn)[0]
	colArray := []string{"area_code", "zip_code", "city_code", "area_name", "name", "short_name", "lng", "lat"}
	dataMap := make(map[string]string)
	for i := 0; i < len(colArray); i++ {
		itemCol := make([]string, 2)
		colName := colArray[i]
		colValue := colMap[colName]
		itemCol[0] = colName
		itemCol[1] = colValue
		dataMap[colName] = colValue
	}
	return dataMap
}
