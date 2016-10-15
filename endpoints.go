package main

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetAllCasesEndpoint(svc CaseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAllCases(ctx)
		if err != nil {
			return getAllCasesResponse{v, err.Error()}, nil
		}
		return getAllCasesResponse{v, ""}, nil
	}
}

func makeGetCaseByIDEndpoint(svc CaseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCaseByIDRequest)
		v, err := svc.GetCaseByID(ctx, req.ID)
		if err != nil {
			return getCaseByIDResponse{v, err.Error()}, nil
		}
		return getCaseByIDResponse{v, ""}, nil
	}
}

func makeGetCasesByStatusNameEndpoint(svc CaseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCasesByStatusNameRequest)
		v, err := svc.GetCasesByStatusName(ctx, req.StatusName)
		if err != nil {
			return getCasesByStatusNameResponse{v, err.Error()}, nil
		}
		return getCasesByStatusNameResponse{v, ""}, nil
	}
}
