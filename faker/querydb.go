package faker

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func getDbPath() string {
	osType := runtime.GOOS
	var dbPath string
	dbPath, _ = os.Getwd()
	if osType == "windows" {
		dbPath = dbPath + "\\..\\faker\\data\\data.db"
	} else if osType == "linux" {
		dbPath = dbPath + "/../faker/data/data.db"
	}
	return dbPath
}

// 返回数据库连接, 若成功第二个参数为空字符串, 否则第二个参数为报错信息
func ConnectSqlite(dbPath string) (*sql.DB, error) {

	//	dbPath := `../faker/data/data.db`
	//	dbPath := getDbPath()
	if _, err := os.Stat(dbPath); err != nil {
		return nil, err
	}
	Conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		defer Conn.Close()
		return nil, err
	}
	if err = Conn.Ping(); err != nil {
		defer Conn.Close()
		return nil, err
	}
	return Conn, nil
}

// 执行查询SQL, 返回查询结果
func QuerySqlite(conn *sql.DB, querySql string, args ...interface{}) ([]map[string]string, error) {
	stmtOut, err := conn.Prepare(querySql)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		log.Fatal(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}

func CreateConn(dbPath string) *sql.DB {
	Conn, err := ConnectSqlite(dbPath)
	if err != nil {
		fmt.Println("连接SQLite错误!!!")
		defer Conn.Close()
		panic(err)
	}
	return Conn
}
