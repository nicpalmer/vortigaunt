package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getter(){
	// Setup mechanism to go and grab the IOCs/IPs
	var m = make(map[string]string)
	m["mld-list.txt"] = "http://www.malwaredomainlist.com/hostslist/ip.txt"
	m["zeus-blocklist.txt"] = "https://zeustracker.abuse.ch/blocklist.php?download=ipblocklist"
	m["Tor-list-all.txt"] = "http://torstatus.blutmagie.de/ip_list_exit.php/Tor_ip_list_ALL.csv"
	m["Multi-Proxy-List.txt"] = "http://multiproxy.org/txt_all/proxy.txt"
	m["ci-badguys.txt"] = "http://cinsscore.com/list/ci-badguys.txt"
	m["blocklist-de.txt"] = "https://lists.blocklist.de/lists/all.txt"
	m["neo23x0.txt"] = "https://raw.githubusercontent.com/Neo23x0/signature-base/master/iocs/hash-iocs.txt"
	for k,v  := range m {
		fmt.Println(k, v)
		resp, err := http.Get(v)
		if err != nil {
			log.Fatalln(err)
		}
		// Create file
		out, err := os.Create(k)
		if err != nil {
			log.Fatalln(err)
		}
		defer out.Close()
		// write to file
		io.Copy(out, resp.Body)
		file, err := os.Open(k)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}