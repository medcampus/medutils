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
	once sync.Once
	dbName, collectionName string
)

type Handler struct {
	C *mgo.Collection
}

// create root DB Session
func createDBSession() (session *mgo.Session, err error) {
	mongoUrl := viper.GetString("mongo.host")
	session, err = mgo.Dial(mongoUrl))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error creating root mongo session %s\n", err.Error()))
	}
	log.Infof("Created ROOT Database Session from %s", mongoUrl)
	dbName = viper.GetString("mongo.dbName")
	collectionName = viper.GetString("mongo.collectionName")
	return session, nil
}

// copy database session
func GetMongoSession() *Handler {

	once.Do(func() {
		var err error
		session, err = createDBSession()
		if err != nil {
			log.Fatalf("Unable to create ROOT database session, Error %v", err)
		}
	})

	return &Handler{session.Copy().DB(dbName).C(collectionName)}
}