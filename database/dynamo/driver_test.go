package dynamo

import (
	"context"
	hlp "github.com/ameniGa/timeTracker/helpers"
	ctxUtl "github.com/ameniGa/timeTracker/helpers/context"
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
