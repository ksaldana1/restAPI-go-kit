package main

import (
	"caseDB"
	"errors"

	"golang.org/x/net/context"
)

var (
	statusNameErr = errors.New("Invalid Status Name")
)

type CaseService interface {
	GetAllCases(ctx context.Context) ([]caseDB.DBCase, error)
	GetCaseByID(ctx context.Context, id string) (caseDB.DBCase, error)
	GetCasesByStatusName(ctx context.Context, statusName string) ([]caseDB.DBCase, error)
}

type caseService struct{}

func (caseService) GetAllCases(ctx context.Context) ([]caseDB.DBCase, error) {
	c, err := caseDB.GetAllCases()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (caseService) GetCaseByID(ctx context.Context, id string) (caseDB.DBCase, error) {
	c, err := caseDB.GetCaseByID(id)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (caseService) GetCasesByStatusName(ctx context.Context, statusName string) ([]caseDB.DBCase, error) {
	statusID, err := caseDB.StatusNameToStatusID(statusName)
	if err != nil {
		return nil, err
	}
	cc, err := caseDB.GetCasesByStatusID(statusID)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
