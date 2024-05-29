package api

import (
	"01-anti-bruteforce/models"
	"01-anti-bruteforce/modules/api/handlers"
	"01-anti-bruteforce/modules/config"
	"01-anti-bruteforce/modules/db"
	"01-anti-bruteforce/modules/ipcheck"
	"01-anti-bruteforce/modules/limiter"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var HOST = os.Getenv("HOST")
var PORT = 5433
var USER = os.Getenv("USER")
var POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var POSTGRES_DBNAME = os.Getenv("POSTGRES_DBNAME")

var (
	N, K, M int
)

func Init() {
	N, M, K = config.Parse()

	db := db.DbConnect(HOST, PORT, USER, POSTGRES_PASSWORD, POSTGRES_DBNAME)
	defer db.Close()

	var Newlimiter = limiter.NewRateLimiter()
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var request models.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		request_ip := ipcheck.ParseIP(r)

		if ipcheck.CheckIpinWL(request_ip, db) {
			log.Println("IP at whitelist")
			result := handlers.Login(request.Username, request.Password)
			fmt.Fprint(w, result)
		} else if ipcheck.CheckIpinBL(request_ip, db) {
			http.Error(w, "Sorry, your IP at blacklist", http.StatusBadRequest)
		} else {
			log.Println("default")
			if Newlimiter.IsAllowed(request.Username, request.Password, request_ip, N, M, K) {
				result := handlers.Login(request.Username, request.Password)
				fmt.Fprint(w, result)
			} else {
				http.Error(w, "Authorization blocked due to rate limit", http.StatusBadRequest)
			}
			time.Sleep(time.Millisecond * 50)
		}
	})

	http.HandleFunc("/bl-add", func(w http.ResponseWriter, r *http.Request) {
		var bl_add models.ListRequest
		err := json.NewDecoder(r.Body).Decode(&bl_add)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		status_add_bl := handlers.AddIPBlackList(bl_add.SUBNET, db)
		fmt.Fprint(w, status_add_bl)
	})

	http.HandleFunc("/bl-delete", func(w http.ResponseWriter, r *http.Request) {
		var bl_delete models.ListRequest
		err := json.NewDecoder(r.Body).Decode(&bl_delete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		status_del_bl := handlers.DelIPBlackList(bl_delete.SUBNET, db)
		fmt.Fprint(w, status_del_bl)
	})

	http.HandleFunc("/wl-add", func(w http.ResponseWriter, r *http.Request) {
		var wl_add models.ListRequest
		err := json.NewDecoder(r.Body).Decode(&wl_add)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		status_add_bl := handlers.AddIPWhiteList(wl_add.SUBNET, db)
		fmt.Fprint(w, status_add_bl)
	})

	http.HandleFunc("/wl-delete", func(w http.ResponseWriter, r *http.Request) {
		var wl_delete models.ListRequest
		err := json.NewDecoder(r.Body).Decode(&wl_delete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		status_del_bl := handlers.DelIPWhiteList(wl_delete.SUBNET, db)
		fmt.Fprint(w, status_del_bl)
	})

	http.HandleFunc("/clear-bucket", func(w http.ResponseWriter, r *http.Request) {
		var clear_bucket models.ClearBucket
		err := json.NewDecoder(r.Body).Decode(&clear_bucket)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		status_clear_bucket := Newlimiter.ClearBucket(Newlimiter, clear_bucket.Login, clear_bucket.Ip)
		fmt.Fprint(w, status_clear_bucket)
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "200 OK")

	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "200 OK")

	})
	log.Println("Api server starting on 5001 port...")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Println("There was an error listening on port :5001", err)
	}

}
