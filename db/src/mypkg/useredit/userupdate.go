package useredit

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	//godotenv.Load(".env")
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

func Userupdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		log.Println(body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		keyVal1 := make(map[string]string)
		json.Unmarshal(body, &keyVal1)
		nameid := keyVal1["nameid"]
		name := keyVal1["name"]

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//TODO pointの制約
		
		mysqlUser := os.Getenv("mysqlUser")
		mysqlUserPwd := os.Getenv("mysqlUserPwd")
		mysqlDatabase := os.Getenv("mysqlDatabase")
		userPasswordDbname := mysqlUser + ":" + mysqlUserPwd + "@/" + mysqlDatabase
		db, err := sql.Open("mysql", userPasswordDbname)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(body)
			log.Println(err)
			return
		}
		defer db.Close()

		_, err = db.Exec("UPDATE name_list SET name=? WHERE nameid=?", name, nameid)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
