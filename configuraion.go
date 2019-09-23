package main

import (
    "os"
    "fmt"
    "encoding/json"
)

type Configuration struct {
    ConnectionString string `json:"rabbit_connection_string"`
    QueueName string `json:"rabbit_endpoint_queue"`
    LogEmailAppName string `json:"log_app_name"`
    LogEmailHost string `json:"log_email_smtp_host"`
    LogEmailPort int `json:"log_email_smtp_port"`
    LogEmailFrom string `json:"log_email_smtp_from"`
    LogEmailTo string `json:"log_email_smtp_to"`
    LogEmailUser string `json:"log_email_smtp_username"`
    LogEmailPassword string `json:"log_email_smtp_passwd"`
    WebServer string `json:"webhook_server"`
}

func configuration() Configuration {
    file, Err := os.Open("config.json")
    if Err != nil {
        fmt.Println("error while loading config from config.json :", Err)
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }

    return configuration
}