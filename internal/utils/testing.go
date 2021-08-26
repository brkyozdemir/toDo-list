package utils

import (
	"database/sql"
	"fmt"
)

func RefreshDatabase() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + "test_db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE todos;")
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("CREATE TABLE todos ( id bigint unsigned auto_increment primary key, description longtext, created_at datetime(3), updated_at datetime(3), deleted_at datetime(3))")
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTestDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + "test_db")
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("USE " + "test_db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE todos ( id bigint unsigned auto_increment primary key, description longtext, created_at datetime(3), updated_at datetime(3), deleted_at datetime(3))")
	if err != nil {
		fmt.Println(err)
	}
}
