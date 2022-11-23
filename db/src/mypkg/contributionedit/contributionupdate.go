package contributionedit

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Contributionupdate(w http.ResponseWriter, r *http.Request) {
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
		id := keyVal1["id"]
		contributorId := keyVal1["contributorId"]
		point, err := strconv.Atoi(keyVal1["point"])
		message := keyVal1["message"]

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//TODO pointの制約

		//godotenv.Load(".env")
		//mysqlUser := os.Getenv("mysqlUser")
		//mysqlUserPwd := os.Getenv("mysqlUserPwd")
		//mysqlDatabase := os.Getenv("mysqlDatabase")
		//userPasswordDbname := mysqlUser + ":" + mysqlUserPwd + "@/" + mysqlDatabase
		//db, err := sql.Open("mysql", userPasswordDbname)
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//defer db.Close()

		_, err = db.Exec("UPDATE contribution_list SET contributorid=?, point=?, message=? WHERE id=?", contributorId, strconv.Itoa(point), message, id)
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
