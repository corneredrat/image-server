package notify
import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/corneredrat/image-server/api-server/config"
)



func Notify(message map[string]string) err {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}
	defer p.Close()

}