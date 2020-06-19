package dynamo

import (
	"context"
	"errors"
	"github.com/ameniGa/timeTracker/config"
	mdl "github.com/ameniGa/timeTracker/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sethvargo/go-password/password"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Repo struct {
	conf *config.Presence
	timeout time.Duration
	*dynamodb.DynamoDB
	log *logrus.Logger
}

// NewPresenceHandler creates a dynamo client and returns a Repo
func NewPresenceHandler(dbConf *config.Presence,timeout time.Duration, logger *logrus.Logger) *Repo {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoClient := dynamodb.New(sess)
	dynamoHandler := Repo{
		conf:     dbConf,
		timeout: timeout,
		DynamoDB: dynamoClient,
		log:      logger,
	}
	return &dynamoHandler
}

// DbAddRating creates a rating for the given provider and with the given information and save it in the database
func (repo *Repo) DbAddUser(ctx context.Context, userID, userName string, ch chan<- error) {
	defer close(ch)
	if ok, err := IsValidCtx(ctx, repo.timeout); !ok {
		repo.log.Error(err)
		ch <- errors.New("invalid input")
		return
	}
	if !IsValidFields(userID, userName) {
		repo.log.Error("invalid input")
		ch <- errors.New("invalid input")
		return
	}

	user, err := prepareUserModel(userID, userName)
	if err != nil {
		repo.log.Error(err)
		ch <- err
		return
	}

	condition := "attribute_not_exists(ServiceName)"
	userMarshalled, _ := dynamodbattribute.MarshalMap(user)
	input := dynamodb.PutItemInput{
		ConditionExpression: &condition,
		Item:                userMarshalled,
		TableName:           aws.String(repo.conf.TableName),
	}
	_, err = repo.PutItemWithContext(ctx, &input)
	if err != nil {
		repo.log.Error(err)
		ch <- err
		return
	}
	ch <- nil
	return
}

func IsValidCtx(ctx context.Context, maxTimeout time.Duration) (bool, error){
	// check timeout and sets default timeout if not specified
	if _, ok := ctx.Deadline(); !ok {
		return false, errors.New("missing context deadline")
	}
	if deadline, _ := ctx.Deadline(); deadline.Sub(time.Now()) > maxTimeout*time.Second {
		return false, errors.New("context deadline exceeded")
	}
	return true, nil
}

func prepareUserModel(userID, userName string) (*mdl.User, error) {
	pass, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		return nil, err
	}
	user := &mdl.User{
		UserID:    userID,
		UserName:  userName,
		Password:  pass,
		CreatedAt: uint64(time.Now().Unix()),
	}
	return user, nil
}

func IsValidFields(id string, name string) bool {
	return !isStringEmpty(id) && !isStringEmpty(name)
}

// isStringEmpty checks if a string is empty, returns bool
func isStringEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}
