package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID  int    `db:"id"`
	NUM string `db:"studentNum"`
	KEY string `db:"secretKey"`
}

type Userlist []User

func main() {
	// var userlist Userlist

	db, err := sqlx.Open("mysql", "ie-user:nagayan@tcp(10.0.0.58:3306)/sshr")
	if err != nil {
		log.Fatal(err)
	}

	// user := []User{}
	// row := db.QueryRow("select * from StudentInfo where studentNum=?", "k198576")
	// var studentNum string
	// err = row.Scan(&studentNum)
	//	err = db.Select(&user, "select * from StudentInfo where studentNum=?", "k198577")
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	// hoge, err := db.Prepare(`select * from StudentInfo where studentNum=?`)
	// if err = hoge.QueryRow("k198576"); != nil {
	// 	return err
	// }

	fmt.Print("hoge")
	// err = db.Get(&hoge, "select * from StudentInfo where studentNum=$1", "hoge")
	// fmt.Printf("%#v\n", hoge)
	// rows, err := db.Queryx("select * from StudentInfo whehe studentNum=$1", "k198576")
	// info := `insert into StudentInfo (studentNum, secretKey) values (?,?)`
	// db.MustExec(info, "hoge", "hogehgjeiruahgiuerwhaihoge")
	// rows, err := db.NamedQuery("select * from StudentInfo whehe studentNum=$1", "k198576")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var user User

	// for rows.Next() {
	// 	err := rows.StructScan(&user)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	userlist = append(userlist, user)
	// }

	// fmt.Println(userlist)

}
