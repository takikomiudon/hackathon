package main

import (
	"database/sql"
	"db/mypkg"
	"db/mypkg/contributionedit"
	"db/mypkg/contributionlist"
	"db/mypkg/ranking"
	"db/mypkg/useredit"
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
	http.HandleFunc("/alltimeranking", ranking.Alltimerank)
	http.HandleFunc("/annualranking", ranking.Annualrank)
	http.HandleFunc("/monthlyranking", ranking.Monthlyrank)
	http.HandleFunc("/weeklyranking", ranking.Weeklyrank)
	http.HandleFunc("/dailyranking", ranking.Dailyrank)
	http.HandleFunc("/mycontribution", contributionlist.Mycontribution)
	http.HandleFunc("/mycontributed", contributionlist.Mycontributed)
	http.HandleFunc("/register", useredit.Register)
	http.HandleFunc("/login", mypkg.Login)
	http.HandleFunc("/contributionpost", contributionedit.Contributionpost)
	http.HandleFunc("/contributiondelete", contributionedit.Contributiondelete)
	http.HandleFunc("/contributionupdate", contributionedit.Contributionupdate)
	http.HandleFunc("/home", mypkg.Home)
	http.HandleFunc("/userdelete", useredit.Userdelete)
	http.HandleFunc("/userupdate", useredit.Userupdate)
	http.HandleFunc("/allcontribution", contributionlist.Allcontribution)
	http.HandleFunc("/contributorlist", mypkg.ContributorList)

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
