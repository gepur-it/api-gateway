package main

import (
    "github.com/streadway/amqp"
	"net/http"
	"io/ioutil"
)

type Writer struct {
    writeChan chan []byte
    connection *amqp.Connection
    channel *amqp.Channel
    queueName string
}

func writer(Config Configuration) *Writer {
    conn, err := amqp.Dial(Config.ConnectionString)
    failOnError(err, "Failed to connect to RabbitMQ")
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    return &Writer{
    		writeChan: make(chan []byte),
    		connection: conn,
    		channel:  ch,
    		queueName: Config.QueueName,
    	}
}

func (wrtr *Writer) send(message []byte) {
    err := wrtr.channel.Publish(
        wrtr.queueName,     // exchange
        wrtr.queueName, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "application/json",
            Body:        message,
        },
    )
    failOnError(err, "Failed to publish a message")
}

func (wrtr *Writer) close() {
    defer wrtr.connection.Close()
    defer wrtr.channel.Close()
}

func (wrtr *Writer) webHook(w http.ResponseWriter, r *http.Request) {
    buf, err := ioutil.ReadAll(r.Body)
    failOnError(err, "Failed to publish a message")
    if(len(buf) == 0) {
        notify("got an empty request")
        return
    }
	wrtr.send(buf)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"status\":\"OK\"}"))
}