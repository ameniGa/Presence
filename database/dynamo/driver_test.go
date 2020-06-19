package dynamo

import (
	"context"
	hlp "github.com/ameniGa/timeTracker/helpers"
	ctxUtl "github.com/ameniGa/timeTracker/helpers/context"
	mdl "github.com/ameniGa/timeTracker/models"
	td "github.com/ameniGa/timeTracker/testData/database"
	"testing"
)

func TestRepo_DbAddUser(t *testing.T) {
	logger := hlp.GetLogger()
	tableTest := td.CreateTTAddUser()
	for _, testCase := range tableTest {
		t.Run(testCase.Name, func(t *testing.T) {
			ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
			defer cancel()
			ch := make(chan error, 1)
			db := NewPresenceHandler(testCase.DBConf, testCase.Timeout, logger)
			db.DbAddUser(ctx, testCase.UserID, testCase.Username, ch)
			err := <-ch
			if err == nil && testCase.HasError {
				t.Errorf("expected failure got: %v", err)
			}
			if err != nil && !testCase.HasError {
				t.Errorf("expected success got error: %v", err)
			}
		})
	}
}


func TestRepo_DbAddEntry(t *testing.T) {
	logger := hlp.GetLogger()
	testDataIn, _ := td.CreateTTAddInOut()
	for _, testCase := range testDataIn {
		t.Run(testCase.Name, func(t *testing.T) {
			ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
			defer cancel()
			ch := make(chan error, 1)
			db := NewPresenceHandler(testCase.DBConf, testCase.Timeout, logger)
			db.DbAddEntry(ctx, testCase.UserID, ch)
			err := <-ch
			if err == nil && testCase.HasError {
				t.Errorf("expected failure got: %v", err)
			}
			if err != nil && !testCase.HasError {
				t.Errorf("expected success got error: %v", err)
			}
		})
	}
}


func TestRepo_DbAddExit(t *testing.T) {
	logger := hlp.GetLogger()
	_, testDataOut := td.CreateTTAddInOut()
	for _, testCase := range testDataOut {
		t.Run(testCase.Name, func(t *testing.T) {
			ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
			defer cancel()
			ch := make(chan error, 1)
			db := NewPresenceHandler(testCase.DBConf, testCase.Timeout, logger)
			db.DbAddExit(ctx, testCase.UserID, ch)
			err := <-ch
			if err == nil && testCase.HasError {
				t.Errorf("expected failure got: %v", err)
			}
			if err != nil && !testCase.HasError {
				t.Errorf("expected success got error: %v", err)
			}
		})
	}
}

func TestRepo_DbGetUserByID(t *testing.T) {
	// insert
	logger := hlp.GetLogger()
	testData := td.CreateTTGetUser()
	for _, testCase := range testData {
		t.Run(testCase.Name, func(t *testing.T) {
			ctx, cancel := ctxUtl.AddTimeoutToCtx(context.Background(), 5)
			defer cancel()
			ch := make(chan mdl.UserWithError, 1)
			db := NewPresenceHandler(testCase.DBConf, testCase.Timeout, logger)
			db.DbGetUserByID(ctx, testCase.UserID, ch)
			output := <-ch
			if output.Error == nil && testCase.HasError {
				t.Errorf("expected failure got: %v", output.Error)
			}
			if output.Error != nil && !testCase.HasError {
				t.Errorf("expected success got error: %v", output.Error)
			}
			if output.UserInfo.UserName != testCase.ExpectedName &&  !testCase.HasError {
				t.Errorf("expected correct username : %v got error: %v", testCase.Name, output.UserInfo.UserName)
			}
		})
	}
}