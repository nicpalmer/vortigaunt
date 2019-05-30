# Vortigaunt

Vortigaunt is a friendly way of ingesting open threat intel feeds into Elasticsearch. 

I wanted to write a bit of Go, and figured this was a good exercise.


## Current feeds

Current Vortigaunt grabs the following feeds; 

```bash
"http://www.malwaredomainlist.com/hostslist/ip.txt"
"https://zeustracker.abuse.ch/blocklist.php?download=ipblocklist"
"http://torstatus.blutmagie.de/ip_list_exit.php/Tor_ip_list_ALL.csv"
"http://multiproxy.org/txt_all/proxy.txt"
"http://cinsscore.com/list/ci-badguys.txt"
"https://lists.blocklist.de/lists/all.txt"
"https://raw.githubusercontent.com/Neo23x0/signature-base/master/iocs/hash-iocs.txt"
```

If you own one of these feeds and don't wish for it to be aggregated here, please raise an issue. 

If you wish to add additional feeds, raise an issue and I'll try and add it in. 
## Installation

```bash
git clone https://github.com/nicpalmer/vortigaunt
cd vortigaunt
go build
```

## Usage

Vortigaunt is configured with environment variables, as a bare minimum you'll need to set, `VORTIGAUNT_PORT`, `VORTIGAUNT_HOSTNAME` and `VORTIGAUNT_SCHEME`

```bash
 export VORTIGAUNT_PORT="9200"
 export VORTIGAUNT_HOSTNAME="localhost"
 export VORTIGAUNT_SCHEME="http"
 export VORTIGAUNT_USERNAME="elastic"
 export VORTIGAUNT_PASSWORD="changeme"

./vortigaunt all
```
> Please don't add the scheme to the hostname value, as Vortigaunt will trip up and think the host is something like `https://http://localhost:9200` I'll fix this in a PR soon.
## Contributing

Go is not my primary language so there's probably a ton of bugs here, please file an issue if you run into issues. 

## Thanks

Thanks to all the people out there who contribute to Open Source, especially those who run open threat feeds. 

## License

[MIT](https://choosealicense.com/licenses/mit/)
