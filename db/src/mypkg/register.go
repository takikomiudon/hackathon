package mypkg

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		keyVal1 := make(map[string]string)
		json.Unmarshal(body, &keyVal1)
		name := keyVal1["name"]

		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		var nameid = ulid.MustNew(ulid.Timestamp(t), entropy)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		godotenv.Load(".env")
		mysqlUser := os.Getenv("mysqlUser")
		mysqlUserPwd := os.Getenv("mysqlUserPwd")
		mysqlHost := os.Getenv("mysqlUserHost")
		mysqlDatabase := os.Getenv("mysqlDatabase")
		userPasswordDbname := mysqlUser + ":" + mysqlUserPwd + "@" + mysqlHost + "/" + mysqlDatabase
		db, err := sql.Open("mysql", userPasswordDbname)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec("INSERT INTO name_list (nameid, name, deleted_at) VALUES(?, ?, false)", nameid.String(), name)
		if err != nil {
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
