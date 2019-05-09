package medutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
		"github.com/confluentinc/confluent-kafka-go/kafka"

	"gopkg.in/mgo.v2"
	"sync"
)

var (
	kafkaconn *kafka.Producer
	once    sync.Once
)

func createKafkaConnection() (kafkaconn *kafka.Producer, err error) {
	
	kafkaHostsUrl := viper.GetString("kafkaHost.host")
	kafkaconn, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaHostsUrl})

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error creating root mongo session %s\n", err.Error()))
	}
	log.Infof("Created Kafka Connection from %s", kafkaHostsUrl)

	return session, nil
}


func GetKafkaSession() *kafka.Producer {

	once.Do(func() {
		var err error
		session, err = createKafkaConnection()
		if err != nil {
			log.Fatalf("Unable to create Kafka session, Error %v", err)
		}
	})

	return session.Copy()
}

func CloseRootSession() {
	session.Close()
}