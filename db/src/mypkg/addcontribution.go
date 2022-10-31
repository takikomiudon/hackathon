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
	"strconv"
	"time"
)

func Addcontribution(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

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
		contributor := keyVal1["contributor"]
		point, err := strconv.Atoi(keyVal1["point"])
		message := keyVal1["message"]
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if name == "" || len(name) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//TODO pointの制約

		godotenv.Load(".env")
		mysqlUser := os.Getenv("mysqlUser")
		mysqlUserPwd := os.Getenv("mysqlUserPwd")
		mysqlDatabase := os.Getenv("mysqlDatabase")
		userPasswordDbname := mysqlUser + ":" + mysqlUserPwd + "@/" + mysqlDatabase
		db, err := sql.Open("mysql", userPasswordDbname)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer db.Close()

		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		var id = ulid.MustNew(ulid.Timestamp(t), entropy)

		_, err = db.Exec("INSERT INTO user (id, name, contributor, point, message, time) VALUES(?, ?, ?, ?, ?, ?)", id.String(), name, contributor, strconv.Itoa(point), message, t)
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
