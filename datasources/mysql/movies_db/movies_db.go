package movies_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_movies_username = "mysql_movies_username"
	mysql_movies_password = "mysql_movies_password"
	mysql_movies_host     = "mysql_movies_host"
	mysql_movies_schema   = "mysql_movies_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_movies_username)
	password = os.Getenv(mysql_movies_password)
	host     = os.Getenv(mysql_movies_host)
	schema   = os.Getenv(mysql_movies_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

	fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

}
