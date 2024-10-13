package consumer

import (
	"github.com/ricky97gr/homeOnline/internal/consumer/es"
	"github.com/ricky97gr/homeOnline/internal/consumer/mongodb"
)

func Start() {
	mongodb.InitConsumer("operation", "first", "127.0.0.1:4161")
	es.InitConsumer("operation", "second", "127.0.0.1:4161")
}
