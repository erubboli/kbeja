package workers

import (
	"log"

	"github.com/erubboli/kbeja/metrics"
)

type Performer func(metrics.Metric)

func Execute(performer Performer, name string) {
	forever := make(chan bool)

	go func() {
		msgs := Rabbit()
		for d := range msgs {
			log.Printf("%s received a message: %s", name, d.Body)
			m, err := parseMessage(d)
			if err != nil {
				log.Printf("Error %s", err)
			} else {
				performer(m)
			}
			d.Ack(false)
		}
	}()

	log.Printf("%s: Waiting for messages.", name)
	<-forever
}
