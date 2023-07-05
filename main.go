package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sellerApp/errors"
	"sellerApp/handlers"
	"sellerApp/models"
	"sellerApp/repository"

	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST      = "db"
	PORT      = "3306"
	USER_NAME = "root"
	PASSWORD  = "password"
	DB_NAME   = "SellerApp"
)

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER_NAME, PASSWORD, HOST, PORT, DB_NAME)
	return sql.Open("mysql", connStr)
}

func main() {

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db.Ping())

	repo := repository.New(db)
	orderHandler := handlers.New(*repo)

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			// extract qyeryt param
			var f models.Filter
			f.UserID = r.URL.Query().Get("userId")
			f.ProductID = r.URL.Query().Get("productId")
			f.Price = r.URL.Query().Get("price")
			f.Qty = r.URL.Query().Get("qty")

			resp, err := orderHandler.Get(f)
			if err != nil {
				log.Println(err)
			}
			sendResponse(w, resp, err)
			return

		case http.MethodPost:
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				sendResponse(w, nil, errors.InvalidParam{Param: []string{"body"}})
				return
			}

			var o models.Order
			err = json.Unmarshal(b, &o)
			if err != nil {
				log.Println(err)
				sendResponse(w, nil, errors.InvalidParam{Param: []string{"body"}})
				return
			}

			err = orderHandler.Create(o)
			sendResponse(w, nil, err)
			return
		}
	})

	log.Println("listining @ port 8000")
	http.ListenAndServe(":8000", nil)
}

func sendResponse(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err == nil {
		json.NewEncoder(w).Encode(data)
		return
	}

	var statusCode int

	apiErr, ok := err.(errors.APIError)
	if !ok {
		statusCode = http.StatusInternalServerError
		errResp := errors.ErrorResponse{Message: err.Error()}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	statusCode, msg := apiErr.APIError()

	w.WriteHeader(statusCode)

	errResp := errors.ErrorResponse{Message: msg}
	json.NewEncoder(w).Encode(errResp)
}
