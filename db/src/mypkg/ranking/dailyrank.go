package ranking

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func Dailyrank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		rows, err := db.Query("SELECT name, SUM(point) FROM contribution_list JOIN name_list ON contribution_list.contributorid=name_list.nameid WHERE date(time) = date(now()) NOT deleted_at GROUP BY contributorid ORDER BY SUM(point) DESC;")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		contribution := make([]RankingResForHTTPGet, 0)
		for rows.Next() {
			var u RankingResForHTTPGet
			if err := rows.Scan(&u.Name, &u.Point); err != nil {
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
