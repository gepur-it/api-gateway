package main

import (
	"net/http"
)

func main() {
    config := configuration()
    initLogger(config)
    writer :=  writer(config)
    defer writer.close()

	http.HandleFunc("/endpoint", writer.webHook)

    notify("Server started")
    err := http.ListenAndServe(config.WebServer, nil)
    if err != nil {
        failOnError(err, "webhook error")
    }
}
