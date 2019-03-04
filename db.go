package medutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2"
)

// create root DB Session
func CreateDBSession(mongoUrl string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error creating root mongo session %s\n", err.Error()))
	}

	log.Infof("Created ROOT Database Session from %s", mongoUrl)
	return session, nil
}