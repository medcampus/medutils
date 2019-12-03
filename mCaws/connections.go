package mCaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	once sync.Once
	sess *session.Session
)

const AwsRegion = "us-east-1"

func awsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
		Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
	})
	if err != nil {
		log.Fatal("unable to create aws session")
	}

	return sess
}

func GetAwsSession(cfgs ...*aws.Config) *session.Session {
	once.Do(func() {
		sess = awsSession()
	})

	return sess.Copy(cfgs...)
}
