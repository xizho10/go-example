package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

type Sqlite struct {
} //结构体

type DailyLog struct {
	Id    int64
	Title string
	Time  string
	Note  string
	Link  string
}

func (m *Sqlite) WriteLog(title string, note string, link string) {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	now := time.Now().Format("2006-01-02 15:03:04")
	//DROP TABLE IF EXISTS tbl_logs;
	//CREATE TABLE tbl_logs(id INTEGER PRIMARY KEY, title TEXT, time TEXT, note TEXT,link TEXT);
	//INSERT INTO tbl_logs(title, time,note,link) VALUES('Audi','2022-1-1','note','link');
	sts := `
INSERT INTO tbl_logs(title, time,note,link) VALUES(?,?,?,?);
`
	_, err = db.Exec(sts, title, now, note, link)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table cars created")
}
func (m *Sqlite) ReadLog(id string) interface{} {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM tbl_logs")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dailyLog DailyLog
		err = rows.Scan(&dailyLog.Id, &dailyLog.Title, &dailyLog.Time, &dailyLog.Note, &dailyLog.Link)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %s\n", dailyLog.Id, dailyLog)
		idInt, _ := strconv.Atoi(id)
		if dailyLog.Id == int64(idInt) {
			//dailyLogByte, _ := json.Marshal(dailyLog)
			return dailyLog
		}
	}
	return nil
}
func (m *Sqlite) ChangeLog(id string, title string) interface{} {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStatement := `UPDATE tbl_logs SET title = ? WHERE id = ?;`
	//idInt, _ := strconv.Atoi(id)
	result, err := db.Exec(sqlStatement, title, id)
	if err != nil {
		panic(err)
	}
	res, _ := result.RowsAffected()
	fmt.Println(result.RowsAffected())
	return res
}
