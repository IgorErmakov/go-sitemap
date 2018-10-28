package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type product struct {
	id int
	ASIN string
	page_link string
	LastMod string
}

var db *sql.DB

func DbConnect() (*sql.DB, error) {

	cfg := getConfig()


	dsn := cfg["user"] + ":" + cfg["pwd"] + "@/" + cfg["name"]

	db, err := sql.Open("mysql", dsn)

	return db, err
}

func GetProducts() []map[string]string {

	if db == nil {

		dbConn, err := DbConnect()

		db = dbConn

		if err != nil {
			panic("Can not connect to db")
		}
	}

	defer db.Close()


	sql := `SELECT id, ASIN, page_link
		    FROM products`


	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var items []map[string]string

	currentTime := time.Now().UTC()

	timeFormatted := currentTime.Format("2006-01-02")

	for rows.Next() {

		p := product{}

		err := rows.Scan(&p.id, &p.ASIN, &p.page_link)

		p.LastMod = timeFormatted

		if err != nil {

			fmt.Println(err)
			continue
		}

		arrItem := map[string]string{
			"id": string(p.id),
			"Location": "/dp/" + p.ASIN + "/" + p.page_link + ".html",
			"LastMod": p.LastMod,
		}

		items = append(items, arrItem)
	}

	return items
}

func getConfig() map[string]string  {

	cfgPath := os.Getenv("SITEMAP_CFG")

	//".config.json"

	r, err := os.Open(cfgPath)
	//r, err := os.Open("/webserver/go/config.json")

	if (err != nil) {
		panic(err)
	}

	defer r.Close()

	bt, _ := ioutil.ReadAll(r)

	var cfg map[string]string

	json.Unmarshal(bt, &cfg)

	return cfg
}