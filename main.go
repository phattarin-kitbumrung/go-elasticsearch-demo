package main

import (
	"context"
	"log"
	"time"

	"github.com/olivere/elastic"
)

var ElasticClient *elastic.Client

type Log struct {
	Service   string      `json:"service"`
	Message   interface{} `json:"message"`
	TimeStamp string      `json:"timestamp"`
	Level     string      `json:"level"`
}

func connectElasticSearch(url string) error {
	clientConnect, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))

	if err != nil {
		return err
	}
	ElasticClient = clientConnect

	return nil
}

func AddLogToIndex(client *elastic.Client, message interface{}, level string) error {
	body := Log{
		Service:   "myService",
		Message:   message,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		Level:     level,
	}

	_, err := client.Index().
		Index("myIndex").
		Type("log").
		BodyJson(body).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := connectElasticSearch("http://localhost:9200/")
	if err != nil {
		log.Panic(err)
	}

	AddLogToIndex(ElasticClient, "test message", "info")
}
