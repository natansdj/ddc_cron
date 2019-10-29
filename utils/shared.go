package utils

type RabbitConfig struct {
	AMQPConnectionURL string
}

type AddTask struct {
	Number1 int
	Number2 int
}

var RabbitConf = RabbitConfig{
	AMQPConnectionURL: "amqp://dev:secret@172.18.0.253:5672/",
}
