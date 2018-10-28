package post

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
	post_title string
	post_name string
	post_date string
	Location string
	LastMod string
}

var db *sql.DB

func DbConnect() (*sql.DB, error) {

	cfg := getConfig()


	dsn := cfg["user"] + ":" + cfg["pwd"] + "@/" + cfg["name"]

	db, err := sql.Open("mysql", dsn)

	return db, err
}

func GetPosts() []map[string]string {

	if (db == nil) {

		//fmt.Println("conn")

		dbConn, err := DbConnect()

		db = dbConn

		if (err != nil) {
			panic("Can not connect to db")
		}
	}

	defer db.Close()


	sql := `SELECT id, post_title, post_name, post_date
		    FROM z_posts
		    WHERE post_type = 'post'
		    AND post_status = 'publish'`


	rows, err := db.Query(sql)

	if (err != nil) {
		panic(err)
	}

	defer rows.Close()

	var items []map[string]string

	currentTime := time.Now().UTC()

	timeFormatted := currentTime.Format("2006-01-02")

	for (rows.Next()) {

		p := product{}

		err := rows.Scan(&p.id, &p.post_title, &p.post_name, &p.post_date)

		dateStr := p.post_date[0:4] + "/" + p.post_date[5:7] + "/" + p.post_date[8:10]

		p.Location = "/" + dateStr + "/" + p.post_name

		p.LastMod = timeFormatted

		if (err != nil) {
			fmt.Println(err)
			continue
		}

		arrItem := map[string]string{
			"id": string(p.id),
			"post_title": p.post_title,
			"post_name": p.post_name,
			"post_date": p.post_date,
			"Location": p.Location,
			"LastMod": p.LastMod,
		}

		items = append(items, arrItem)
	}

	return items
}

func getConfig() map[string]string  {

	cfgPath := os.Getenv("SITEMAP_CFG")

	if cfgPath == "" {
		panic("export SITEMAP_CFG=PATH_TO_cfg file")
	}

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