package main

import (
	"caseDB"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()
	svc := caseService{}
	defer caseDB.DB.Close()

	r := mux.NewRouter().StrictSlash(true)
	r.Methods("GET").Path("/cases").Handler(httptransport.NewServer(
		ctx,
		makeGetAllCasesEndpoint(svc),
		decodeGetAllCasesRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/cases/{id}").Handler(httptransport.NewServer(
		ctx,
		makeGetCaseByIDEndpoint(svc),
		decodeGetCaseByIDRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/cases/filter/statusName").Handler(httptransport.NewServer(
		ctx,
		makeGetCasesByStatusNameEndpoint(svc),
		decodeGetCasesByStatusNameRequest,
		encodeResponse,
	))

	fmt.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
