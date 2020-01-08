package utils

import "os"

type RabbitConfig struct {
	AMQPConnectionURL string
}

type AddTask struct {
	Number1 int
	Number2 int
}

var RabbitConf RabbitConfig

func InitRabbitConfig() {
	RabbitConf.Init()
}

func (u *RabbitConfig) Init() {
	// Get the connection string from the environment variable
	url := os.Getenv("AMQP_URL")

	//If it doesnt exist, use the default connection string
	if url == "" {
		url = "amqp://dev:secret@172.18.0.253:5672/"
	}

	u.AMQPConnectionURL = url
}
