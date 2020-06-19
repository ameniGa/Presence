package dynamo

import (
	"github.com/ameniGa/timeTracker/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	conf *config.Presence
	*dynamodb.DynamoDB
	log *logrus.Logger
}


// NewPresenceHandler creates a dynamo client and returns a Repo
func NewPresenceHandler(dbConf *config.Presence, logger *logrus.Logger) *Repo {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoClient := dynamodb.New(sess)
	dynamoHandler := Repo{
		conf:     dbConf,
		DynamoDB: dynamoClient,
		log:      logger,
	}
	return &dynamoHandler
}