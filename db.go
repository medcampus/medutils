package medutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	session *mgo.Session
	once    sync.Once
)

// create root DB Session
func createDBSession() (session *mgo.Session, err error) {
	mongoUrl := viper.GetString("mongo.host")

	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error creating root mongo session %s\n", err.Error()))
	}
	log.Infof("Created ROOT Database Session from %s", mongoUrl)

	return session, nil
}

// copy database session
func GetMongoSession() *mgo.Session {

	once.Do(func() {
		var err error
		session, err = createDBSession()
		if err != nil {
			log.Fatalf("Unable to create ROOT database session, Error %v", err)
		}
	})

	return session.Copy()
}

func CloseRootSession() {
	session.Close()
}
