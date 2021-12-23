package receiver

import (
	"context"
	"time"

	"dev.azure.com/carrotins/hdm/hdm-go.git/logging"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	met "github.com/rcrowley/go-metrics"
)

type Receiver interface {
	Start() error
	Stop()
}

type receiver struct {
	conf       Config
	log        logging.Log
	subscriber *kafka.Subscriber
	tlogTopic  string
	done       chan interface{}
}

func New(cfg *Config) (Receiver, error) {
	return nil, nil
}

func (r *receiver) Start() error {
	tlogMessages, err := r.subscriber.Subscribe(context.Background(), r.tlogTopic)
	if err != nil {
		r.log.Errorf("kafka '%s' subscriber failed, %s", r.tlogTopic, err)
		return err
	}

	r.log.Info("Start")

	go func() {
		tlogTimer := met.GetOrRegisterTimer("dtag.tard.process.tlog", met.DefaultRegistry)
		for {
			select {
			case msg := <-tlogMessages:
				t1 := time.Now()
				// r.processEvt(msg)
				tlogTimer.Update(time.Since(t1))
			case <-r.done:
				return
			}
		}
	}()

	return nil
}

func (r *receiver) Stop() {
	r.done <- "done."
	close(r.done)
	r.subscriber.Close()
	r.log.Info("Stop")
}
