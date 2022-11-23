package mypkg

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"net/url"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		m, _ := url.ParseQuery(r.URL.RawQuery)
		nameid := m["nameid"][0]
		if nameid == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rows, err := db.Query(
			"SELECT CONCAT(IFNULL(starttitle,''),' ',name,' ',IFNULL(endtitle,'')) AS titlename FROM (SELECT name, SUM(point) AS point FROM name_list JOIN contribution_list ON name_list.nameid=contribution_list.contributorid WHERE name_list.nameid=? GROUP BY contributorid) list JOIN starttitle_list ON list.point >= starttitle_list.point AND list.point < starttitle_list.point+100 JOIN endtitle_list ON list.point >= endtitle_list.point AND list.point < endtitle_list.point+100;", nameid)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var user string
		for rows.Next() {
			if err := rows.Scan(&user); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)
				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if user == "" {
			rows, err := db.Query("SELECT name FROM name_list WHERE nameid=?;", nameid)
			if err != nil {
				log.Printf("fail: db.Query, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for rows.Next() {
				if err := rows.Scan(&user); err != nil {
					log.Printf("fail: rows.Scan, %v\n", err)
					if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
						log.Printf("fail: rows.Close(), %v\n", err)
					}
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			user = user + "さん"
		}
		bytes, err := json.Marshal(user)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		return

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
