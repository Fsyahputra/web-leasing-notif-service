package IF

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/IBM/sarama"
)

type Kafka struct {
	Brokers []string
	Topics  []string
	Group   string
	config  *sarama.Config
	cg      sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
	name    string
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

func (k *Kafka) Consume(ctx context.Context) error {
	go func() {
		for {
			select {
			case err, ok := <-k.cg.Errors():
				if !ok {
					return
				}
				log.Println("consumer group error:", err.Error())
			case <-ctx.Done():
				return
			}
		}
	}()

	for {

		if err := k.cg.Consume(ctx, k.Topics, k.handler); err != nil {
			log.Println("consume error:", err)
			time.Sleep(time.Second)
		}
		if ctx.Err() != nil {
			return k.cg.Close()
		}
	}
}

func NewKafka(Brokers []string, Topics []app.EventType, Group string, name string) *Kafka {
	defaultName := "consumer"
	var consumerName string
	if name == "" {
		consumerName = defaultName
	}
	consumerName = name
	topicStrs := make([]string, len(Topics))
	for _, t := range Topics {
		topicStrs = append(topicStrs, string(t))
	}
	kafka := &Kafka{
		Brokers: Brokers,
		Topics:  topicStrs,
		config:  nil,
		Group:   Group,
		cg:      nil,
		name:    consumerName,
	}
	kafka.configure()
	kafka.newConsumer()
	return kafka
}

type KafkaManager struct {
	kafkas []*Kafka
}

func (km *KafkaManager) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	var wg sync.WaitGroup
	for _, k := range km.kafkas {
		wg.Go(
			func() {
				if err := k.Consume(ctx); err != nil {
					log.Printf("Kafka %s failed to stop error: %v", k.name, err)
				}
			})
	}
	log.Println("Kafka consumers are running...")
	<-sig
	cancel()
	log.Println("Shutting down Kafka consumers...")
	wg.Wait()
	log.Println("All Kafka consumers have been shut down.")
}
