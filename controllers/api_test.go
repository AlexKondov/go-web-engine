// +build !mgo

package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"

	"basaigbook/data"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	db := &data.DB{}
	if err := db.Open("unit", "test"); err != nil {
		log.Fatal("error while creating mem data ", err)
	}

	api := &API{
		DB:     db,
		Logger: logger,
	}

	rec := httptest.NewRecorder()

	api.ServeHTTP(rec, req)
	return rec
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
