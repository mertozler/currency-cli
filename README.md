# Currency CLI

currency-cli is a cli tool that allows you to fetch the exchange rates you want.

![CurrencyCLI](https://github.com/mertozler/currency-cli/blob/main/example.gif?raw=true)

## Flags

 - -b, --btc       Get BTC Currency
 -    ,--eth       Get ETH Currency
 - -e, --euro      Get Euro Currency
 - -g, --gold      Get GOLD Currency (Gram)
 - -h, --help      help for get
 - -s, --sterlin   Get Sterlin Currency
 - -u, --usd       Get USD Currency

## Example Usage

 - "currency-cli get -u" -> [USD] [2022-09-11 18:57:49] : Buying: 18.2413₺, Sales: 18.2312₺
 - "currency-cli get --btc" -> [BTC] [2022-09-11 18:57:59] : Buying: 21664.8170$, Sales: 21664.8170$
 - "currency-cli get -e" -> [EURO] [2022-09-11 18:58:08] : Buying: 18.3262₺, Sales: 18.3082₺

## Installation
```bash
go install github.com/mertozler/currency-cli@latest
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
