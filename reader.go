package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"os"
	"strings"
	"time"
)



func reader() {
	host, username, password := config()

	type Format struct {
		Source string
		IP string
		Hash string
		UploadedAt time.Time
	}

	cfg := elasticsearch.Config{
		Addresses: []string{host},
		Username: username,
		Password: password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var filenames= []string{"mld-list.txt", "zeus-blocklist.txt", "Tor-list-all.txt", "Multi-Proxy-List.txt", "ci-badguys.txt", "neo23x0.txt", "blocklist-de.txt"}
	for _, arg := range filenames {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		t := time.Now()
		fmt.Println(t)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(string("ti-go-"+t.Format("2006-01-02")))
			if arg == "neo23x0.txt" {
				data := Format{
					Source: file.Name(),
					Hash:     strings.TrimRight(scanner.Text(), "\n"),
					UploadedAt: time.Now(),
				}
				file, _ := json.MarshalIndent(data, "", " ")
				fmt.Println(string(file))
				req := esapi.IndexRequest{
					Index:   string("ti-go-"+t.Format("2006-01-02")),
					Body:    strings.NewReader(string(file)),
					Refresh: "false",
					Pipeline: "hash",
				}
				res, err := req.Do(context.Background(), es)
				if err != nil {
					log.Fatalf("Error getting response: %s", err)
				}
				if res.IsError() {
					log.Printf("[%s]", res.Status())
				} else {
					// Deserialize the response into a map.
					var r map[string]interface{}
					if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
						log.Printf("Error parsing the response body: %s", err)
					} else {
						// Print the response status and indexed document version.
						log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
					}
				}
				defer res.Body.Close()
			} else {
				data := Format{
					Source: file.Name(),
					IP:     strings.TrimRight(scanner.Text(), "\n"),
					UploadedAt: time.Now(),
				}
				file, _ := json.MarshalIndent(data, "", " ")
				fmt.Println(string(file))
				req := esapi.IndexRequest{
					Index:   string("ti-go-"+t.Format("2006-01-02")),
					Body:    strings.NewReader(string(file)),
					Refresh: "false",
					Pipeline: "ip",
				}
				res, err := req.Do(context.Background(), es)
				if err != nil {
					log.Fatalf("Error getting response: %s", err)
				}
				if res.IsError() {
					log.Printf("[%s]", res.Status())
				} else {
					// Deserialize the response into a map.
					var r map[string]interface{}
					if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
						log.Printf("Error parsing the response body: %s", err)
					} else {
						// Print the response status and indexed document version.
						log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
					}
				}
				defer res.Body.Close()
			}


		}
		defer file.Close()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}