package contributionlist

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"net/url"
)

type ContributionResForHTTPGet struct {
	Contributor string
	Point       int
	Message     string
}

func Mycontribution(w http.ResponseWriter, r *http.Request) {
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

		rows, err := db.Query("SELECT name, point, message FROM contribution_list JOIN name_list ON contribution_list.nameid = name_list.nameid WHERE contributorid=?;", nameid)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		contribution := make([]ContributionResForHTTPGet, 0)
		for rows.Next() {
			var u ContributionResForHTTPGet
			if err := rows.Scan(&u.Contributor, &u.Point, &u.Message); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)
				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			contribution = append(contribution, u)
		}
		bytes, err := json.Marshal(contribution)
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
