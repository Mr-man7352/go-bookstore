package main

import (
	"log"
	"net/http"
	"github.com/Mr-man7352/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/muz"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
)