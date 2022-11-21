package main

import (
	"database/sql"
	"db/mypkg"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func main() {
	http.HandleFunc("/alltimeranking", mypkg.Alltimerank)
	http.HandleFunc("/annualranking", mypkg.Annualrank)
	http.HandleFunc("/monthlyranking", mypkg.Monthlyrank)
	http.HandleFunc("/weeklyranking", mypkg.Weeklyrank)
	http.HandleFunc("/dailyranking", mypkg.Dailyrank)
	http.HandleFunc("/mycontribution", mypkg.Mycontribution)
	http.HandleFunc("/mycontributed", mypkg.Mycontributed)
	http.HandleFunc("/register", mypkg.Register)
	http.HandleFunc("/login", mypkg.Login)
	http.HandleFunc("/contributionpost", mypkg.Contributionpost)
	http.HandleFunc("/contributiondelete", mypkg.Contributiondelete)
	http.HandleFunc("/contributionupdate", mypkg.Contributionupdate)
	http.HandleFunc("/home", mypkg.Home)
	http.HandleFunc("/userdelete", mypkg.Userdelete)
	http.HandleFunc("/userupdate", mypkg.Userupdate)
	http.HandleFunc("/allcontribution", mypkg.Allcontribution)

	closeDBWithSysCall()

	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
