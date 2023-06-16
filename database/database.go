package database

import (
	"database/sql"
	"fmt"
)


func Init() (*sql.DB) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	checkErr(err)
  
	  //defer db.Close()
	  // make sure connection is available
	  err = db.Ping()
	  checkErr(err)
	  fmt.Printf("Connection successfully")
  
	  return db
  }
  
  func checkErr(err error) {
	if err != nil {
	  fmt.Print(err.Error())
	}
  }