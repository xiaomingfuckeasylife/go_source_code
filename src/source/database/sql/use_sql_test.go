package sql

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func TestSql(t *testing.T){
	db := getDb()
	r , err := db.Query("select name from test where id  <> ?",0)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer r.Close()
	for r.Next() {
		var name string
		if err = r.Scan(&name) ; err != nil {
			t.Fatal(err.Error())
		}
		fmt.Printf("name is %s \n" ,name )
	}
	if err = r.Err() ; err != nil {
		t.Fatal(err.Error())
	}
}

// we can have multiple result in one result . cool  right .
func TestMultpleResultSet(t *testing.T){
	q := `
		select id , name from test
	`
	db := getDb()
	r , err := db.Query(q)
	defer r.Close()
	if err != nil {
		t.Fatal(err.Error())
	}
	for r.Next() {
		type result struct{
			id int
			name string
		}
		rst := &result{}
		err := r.Scan(&rst.id,&rst.name)
		if err != nil {
			t.Fatal(err.Error())
		}
		println(rst.id, rst.name)
	}
	if r.NextResultSet() {
		for r.Next() {
			type result struct{
				id int
				name string
			}
			rst := &result{}
			err := r.Scan(&rst.id,&rst.name)
			if err != nil {
				t.Fatal(err.Error())
			}
			println(rst.id, rst.name)
		}
	}
}


func getDb() *sql.DB{
	//db , err := sql.Open("mysql","root:123456@/test")
	db , err := sql.Open("mysql","cde:stqySWTYXIDlHms3@(uatdb.c0itmuga9use.rds.cn-north-1.amazonaws.com.cn:3306)/cde")
	//defer  db.Close() // normally we do not need to close the database.
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return db
}







