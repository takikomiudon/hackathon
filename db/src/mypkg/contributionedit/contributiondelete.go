package contributionedit

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

func Contributiondelete(w http.ResponseWriter, r *http.Request) {
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
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("DELETE FROM contribution_list WHERE id=?", id)
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
