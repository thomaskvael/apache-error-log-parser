package parser

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import mysql driver
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "apache_error"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func insertDatabase(logMap map[string]string) {

	/*
		for k, v := range logMap {
			fmt.Println(k, v)
		}
	*/

	// Parse and format time.Time to Mysql Datetime
	t, _ := time.Parse("Mon Jan 02 15:04:05.999999 2006", logMap["time"])
	mysqlDatetime := t.Format("2006-01-02 15:04:05")

	fmt.Println("fddfdf")
	fmt.Println(mysqlDatetime)
	fmt.Println(logMap["loglevel"])
	fmt.Println(logMap["pid"])
	fmt.Println(logMap["tid"])
	fmt.Println(logMap["apr"])
	fmt.Println(logMap["client"])
	fmt.Println(logMap["message"])

	db := dbConn()

	// insert
	stmt, err := db.Prepare("INSERT INTO logs (time, loglevel, pid, tid, apr, client, message) VALUES (?,?,?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(mysqlDatetime, logMap["loglevel"], logMap["pid"], logMap["tid"], logMap["apr"], logMap["client"], logMap["message"])
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	/*
		stmt, err = db.Prepare("update userinfo set username=? where uid=?")
		checkErr(err)

		res, err = stmt.Exec("astaxieupdate", id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	*/

	// query
	/*
		rows, err := db.Query("SELECT * FROM userinfo")
		checkErr(err)

		for rows.Next() {
				var uid int
				var username string
				var department string
				var created string
				err = rows.Scan(&uid, &username, &department, &created)
				checkErr(err)
				fmt.Println(uid)
				fmt.Println(username)
				fmt.Println(department)
				fmt.Println(created)
		}
	*/

	// delete
	/*
			stmt, err = db.Prepare("delete from userinfo where uid=?")
			checkErr(err)


		res, err = stmt.Exec(id)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)


		fmt.Println(affect)
	*/

	//db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
