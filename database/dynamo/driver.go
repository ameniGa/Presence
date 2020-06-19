package dynamo

import (
	"context"
	"errors"
	"github.com/ameniGa/timeTracker/config"
	ctxUtl "github.com/ameniGa/timeTracker/helpers/context"
	vld "github.com/ameniGa/timeTracker/helpers/validators"
	mdl "github.com/ameniGa/timeTracker/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sethvargo/go-password/password"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Repo struct {
	conf    *config.Presence
	timeout time.Duration
	*dynamodb.DynamoDB
	log *logrus.Logger
}

// NewPresenceHandler creates a dynamo client and returns a Repo
func NewPresenceHandler(dbConf *config.Presence, timeout time.Duration, logger *logrus.Logger) *Repo {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoClient := dynamodb.New(sess)
	dynamoHandler := Repo{
		conf:     dbConf,
		timeout:  timeout,
		DynamoDB: dynamoClient,
		log:      logger,
	}
	return &dynamoHandler
}

// DbAddUser creates a user for the given id and with the given username and save it in the database
func (repo *Repo) DbAddUser(ctx context.Context, userID, userName string, ch chan<- error) {
	defer close(ch)
	if ok, err := ctxUtl.IsValidCtx(ctx, repo.timeout); !ok {
		repo.log.Error(err)
		ch <- errors.New("invalid input")
		return
	}
	if !IsValidFields(userID, userName) {
		repo.log.Error("invalid input")
		ch <- errors.New("invalid input")
		return
	}

	if !vld.IsValidID(userID) {
		repo.log.Error("invalid id")
		ch <- errors.New("invalid id")
		return
	}

	user, err := prepareUserModel(userID, userName)
	if err != nil {
		repo.log.Error(err)
		ch <- err
		return
	}
	repo.log.Error(repo.conf.UserTableName)
	condition := "attribute_not_exists(UserID)"
	userMarshalled, _ := dynamodbattribute.MarshalMap(user)
	input := dynamodb.PutItemInput{
		ConditionExpression: &condition,
		Item:                userMarshalled,
		TableName:           aws.String(repo.conf.UserTableName),
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

// DbAddEntry creates a new entry for the given userId save it in the database
func (repo *Repo) DbAddEntry(ctx context.Context, userID string, ch chan<- error) {
	defer close(ch)

	if ok, err := ctxUtl.IsValidCtx(ctx, repo.timeout); !ok {
		repo.log.Error(err)
		ch <- err
		return
	}

	if vld.IsStringEmpty(userID) || !vld.IsValidID(userID) {
		repo.log.Error("invalid input")
		ch <- errors.New("invalid input")
		return
	}

	timeTrack := prepareEntryTimeTrackModel(userID)
	userMarshalled, _ := dynamodbattribute.MarshalMap(timeTrack)
	input := dynamodb.PutItemInput{
		Item:      userMarshalled,
		TableName: aws.String(repo.conf.TimeTableName),
	}
	_, err := repo.PutItemWithContext(ctx, &input)
	if err != nil {
		repo.log.Error(err)
		ch <- err
		return
	}
	ch <- nil
	return
}

// DbAddExit updates a exit time for the given userId save it in the database
func (repo *Repo) DbAddExit(ctx context.Context, userID string, ch chan<- error) {
	defer close(ch)

	if ok, err := ctxUtl.IsValidCtx(ctx, repo.timeout); !ok {
		repo.log.Error(err)
		ch <- err
		return
	}

	if vld.IsStringEmpty(userID) || !vld.IsValidID(userID) {
		repo.log.Error("invalid input")
		ch <- errors.New("invalid input")
		return
	}

	timeTrack := prepareExitTimeTrackModel(userID)

	input := dynamodb.UpdateItemInput{

		ExpressionAttributeNames: map[string]*string{
			"#o": aws.String("Exit"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":o": {
				S: aws.String(timeTrack.Exit),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(timeTrack.UserID),
			},
			"Date": {
				S: aws.String(timeTrack.Date),
			},
		},
		TableName:        aws.String(repo.conf.TimeTableName),
		UpdateExpression: aws.String("set #o = :o"),
	}

	_, err := repo.UpdateItemWithContext(ctx, &input)
	if err != nil {
		repo.log.Error(err)
		ch <- err
		return
	}
	ch <- nil
	return
}

func prepareEntryTimeTrackModel(id string) mdl.TimeTrack {
	time := mdl.TimeTrack{
		UserID: id,
		Entry:  strconv.FormatInt(time.Now().Unix(), 10),
		Date:   time.Now().Format("2006-01-02"),
	}
	return time
}

func prepareExitTimeTrackModel(id string) mdl.TimeTrack {
	time := mdl.TimeTrack{
		UserID: id,
		Exit:   strconv.FormatInt(time.Now().Unix(), 10),
		Date:   time.Now().Format("2006-01-02"),
	}
	return time
}

func prepareUserModel(userID, userName string) (*mdl.User, error) {
	pass, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		return nil, err
	}
	user := &mdl.User{
		UserID:      userID,
		UserName:    userName,
		Password:    pass,
		CreatedAt:   uint64(time.Now().Unix()),
		PassChanged: false,
	}
	return user, nil
}

func IsValidFields(id string, name string) bool {
	return !vld.IsStringEmpty(id) && !vld.IsStringEmpty(name)
}
