package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"os"
)

var (
	res *esapi.Response
	err error
	indexName = "ti-go"
)
func setup(){
	host, username, password := config()
	cfg := elasticsearch.Config{
		Addresses: []string{host},
		Username: username,
		Password: password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// Use IO Reader to load the ingestion pipeline in
	ip, err := os.Open("./pipelines/ti-ip-pipeline.json")
	fmt.Println("Installing the IP Pipeline")
	if err != nil {}
	res, err := es.Ingest.PutPipeline("ip", ip)
	if res.IsError() {
		fmt.Println("error: %s", res)
	}
	hash, err := os.Open("./pipelines/ti-hash-pipeline.json")
	fmt.Println("Installing the Hash Pipeline")
	if err != nil {}
	resp, err := es.Ingest.PutPipeline("hash", hash)
	if resp.IsError() {
		fmt.Println("error: %s", res)
	}
}