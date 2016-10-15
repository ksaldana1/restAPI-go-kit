package main

import (
	"caseDB"
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouter = errors.New("missing path variable")
)

func decodeGetAllCasesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getAllCasesRequest{}, nil
}

func decodeGetCaseByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouter
	}
	return getCaseByIDRequest{ID: id}, nil
}

func decodeGetCasesByStatusNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getCasesByStatusNameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	c = httptransport.SetResponseHeader("Content-Type", "application/json")(c, w)
	return json.NewEncoder(w).Encode(response)
}

type getAllCasesRequest struct{}

type getAllCasesResponse struct {
	Data []caseDB.DBCase `json:"data"`
	Err  string          `json:"err, omitempty"`
}

type getCaseByIDRequest struct {
	ID string
}

type getCaseByIDResponse struct {
	Data caseDB.DBCase `json:"data"`
	Err  string        `json:"err, omitempty"`
}

type getCasesByStatusNameRequest struct {
	StatusName string `json:"statusName"`
}

type getCasesByStatusNameResponse struct {
	Data []caseDB.DBCase `json:"data"`
	Err  string          `json:"err, omitempty"`
}
