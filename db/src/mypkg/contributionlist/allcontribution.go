package contributionlist

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	godotenv.Load(".env")
	mysqlUser := os.Getenv("mysqlUser")
	mysqlUserPwd := os.Getenv("mysqlUserPwd")
	mysqlHost := os.Getenv("mysqlHost")
	mysqlDatabase := os.Getenv("mysqlDatabase")
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

type AllContributionResForHTTPGet struct {
	Name        string
	Contributor string
	Point       int
	Message     string
}

func Allcontribution(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		rows, err := db.Query("SELECT user, name AS contributor, point, message FROM (SELECT name AS user, contributorid, point, message FROM contribution_list JOIN name_list ON name_list.nameid=contribution_list.nameid) list JOIN name_list ON list.contributorid=name_list.nameid;")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		contribution := make([]AllContributionResForHTTPGet, 0)
		for rows.Next() {
			var u AllContributionResForHTTPGet
			if err := rows.Scan(&u.Name, &u.Contributor, &u.Point, &u.Message); err != nil {
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
