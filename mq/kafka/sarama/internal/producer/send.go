package producer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"

	"sarama/internal/config"
)

func init() {
	config.InitConfig()
}

func SendMessages() {
	keepRunning := true
	// 初始化客户端配置
	cfg := config.GetConfig()
	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}
	producerProvider := newProducerProvider(cfg.Brokers, func() *sarama.Config {
		config := sarama.NewConfig()
		config.Version = version
		config.Producer.Idempotent = true
		config.Producer.Return.Errors = false
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
		config.Producer.Transaction.Retry.Backoff = 10
		config.Producer.Transaction.ID = "txn_producer"
		config.Net.MaxOpenRequests = 1
		return config
	})

	// 生产数据
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	for i := 0; i < cfg.Producer.Count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					produceTestRecord(producerProvider)
				}
			}
		}()
	}

	// 监听信号
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		<-sigterm
		log.Println("terminating: via signal")
		keepRunning = false
	}

	// 关闭生产者
	cancel()
	wg.Wait()
	producerProvider.clear()
}
