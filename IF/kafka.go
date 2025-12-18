package IF

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

type Kafka struct {
	Brokers []string
	Topics  []string
	Group   string
	config  *sarama.Config
	cg      sarama.ConsumerGroup
}

func (k *Kafka) configure() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	k.config = config

}

func (k *Kafka) newConsumer() error {
	if k.config == nil {
		k.configure()
	}
	cg, err := sarama.NewConsumerGroup(k.Brokers, k.Group, k.config)
	if err != nil {
		return err
	}
	k.cg = cg
	return nil
}

func (k *Kafka) Consume(handler sarama.ConsumerGroupHandler) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case err, ok := <-k.cg.Errors():
				if !ok {
					return
				}
				log.Println("consumer group error:", err)
			case <-ctx.Done():
				return
			}
		}
	}()
	for {
		select {
		case <-sign:
			log.Println("shutdown signal received")
			cancel()
			return k.cg.Close()

		default:
			if err := k.cg.Consume(ctx, k.Topics, handler); err != nil {
				log.Println("consume error:", err)
				time.Sleep(time.Second)
			}
			if ctx.Err() != nil {
				return k.cg.Close()
			}
		}
	}
}

func NewKafka(Brokers []string, Topics []string, Group string) *Kafka {
	kafka := &Kafka{
		Brokers: Brokers,
		Topics:  Topics,
		config:  nil,
		Group:   Group,
		cg:      nil,
	}
	kafka.configure()
	kafka.newConsumer()
	return kafka
}
