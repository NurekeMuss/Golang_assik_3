package main

import (
	"fmt"
	"log"
	"net/http"

	"clean_code/pkg/store/postgres"
	delivery "clean_code/services/contact/internal/delivery/http"
	repository "clean_code/services/contact/internal/repository/storage/postgres"
	contactUseCase "clean_code/services/contact/internal/useCase/contact"
	groupUseCase "clean_code/services/contact/internal/useCase/group"
)

func main() {
	config := postgres.NewDBConfig("localhost", 5432, "postgres", "12605291", "postgres")
	db, err := postgres.ConnectToDB(config)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := repository.New(db)
	ucGroup := groupUseCase.New(r)
	ucContact := contactUseCase.New(r)

	d := delivery.New(ucContact, ucGroup)

	addr := 4000
	addrStr := fmt.Sprintf(":%d", addr)

	log.Printf("Starting server on port: %d", addr)

	if err := http.ListenAndServe(addrStr, d.Router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
