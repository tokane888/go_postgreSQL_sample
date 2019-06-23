package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string
}

var Db *sql.DB

func Init() {
	var err error
	Db, err = sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}

	// sql.Open()はConnection poolを作成するだけで実際には接続していないため認証などではエラーにならない。
	// Pingを打つことで設定値等検証可能
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, name from public.user where id = $1", id).Scan(&user.Id, &user.Name)
	fmt.Println(user.Id, user.Name)
	return
}

func main() {
	Init()
	GetPost(1)
	Db.Close()
}
