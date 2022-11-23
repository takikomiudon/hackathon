package contributionedit

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
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

func Contributionpost(w http.ResponseWriter, r *http.Request) {
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
		nameid := keyVal1["nameid"]
		contributorId := keyVal1["contributorId"]
		point, err := strconv.Atoi(keyVal1["point"])
		message := keyVal1["message"]

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}

		t := time.Now().In(jst)
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		var id = ulid.MustNew(ulid.Timestamp(t), entropy)

		_, err = db.Exec("INSERT INTO contribution_list (id, nameid, contributorid, point, message, time) VALUES(?, ?, ?, ?, ?, ?)", id.String(), nameid, contributorId, strconv.Itoa(point), message, t)
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
