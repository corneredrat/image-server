package notifier
import (
	"github.com/segmentio/kafka-go"
	"github.com/corneredrat/image-server/api-server/config"
	log "github.com/sirupsen/logrus"
	"context"
	"encoding/json"
	"time"
	"fmt"
)

const (
	CreateOp = "createOperation"
	DeleteOp = "deleteOperation"
)

func Notify(message map[string]string, topic string)  {
	defer func() {
        if r := recover(); r != nil {
            log.Error("notifier panicked, ...recovered")
        }
    }()
	partition := 0
	msg := fmt.Sprintf("attempting to send notifification to Kafka at : %v:%v",config.Options.Kafka.URL,config.Options.Kafka.PORT)
	log.Info(msg)
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.Options.Kafka.URL+":"+config.Options.Kafka.PORT, topic, partition)
	if err != nil {
		log.Error("failed to dial leader:", err)
	}
	defer conn.Close()
	byteMessage, _ := json.Marshal(message)
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: byteMessage},
	)
	if err != nil {
		log.Error("failed to write messages:", err)
	}
}