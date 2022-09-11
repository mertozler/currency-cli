/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mertozler/currency-cli/exchange"
	"github.com/mertozler/currency-cli/model"
	"github.com/spf13/cobra"
	"time"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "This CLI brings up instant exchange rates",
	Long: `

For use, you need to enter a command like currency-cli get -'currency-name'.
For example, you should use "currency-cli get -u" to access USD information or "currency-cli get --euro" to access EURO information
If you want to see the currently existing currencies on the CLI, you can access the list with the "currency-cli get -h" command.

Note: If the API is not available, the exchange information will not be accessible. If you observe this, you can inform meozler@gmail.com.`,
	Run: func(cmd *cobra.Command, args []string) {
		isUsd, _ := cmd.Flags().GetBool("usd")
		isEuro, _ := cmd.Flags().GetBool("euro")
		isSterlin, _ := cmd.Flags().GetBool("sterlin")
		isBtc, _ := cmd.Flags().GetBool("btc")
		isEth, _ := cmd.Flags().GetBool("eth")
		isGold, _ := cmd.Flags().GetBool("gold")

		if isUsd {
			getUsd()
		} else if isEuro {
			getEuro()
		} else if isSterlin {
			getSterlin()
		} else if isBtc {
			getBtc()
		} else if isEth {
			getEth()
		} else if isGold {
			getGold()
		} else {
			fmt.Println("You have to choose the currency type, you can use the \"currency-cli get -h\" command to learn the flags.")
		}
	},
}

func getUsd() {
	exchange := getExchange()
	printCurrency("USD", "₺", exchange.Usd.Satis, exchange.Usd.Alis)
}

func getEuro() {
	exchange := getExchange()
	printCurrency("EURO", "₺", exchange.Eur.Satis, exchange.Eur.Alis)
}
func getBtc() {
	exchange := getExchange()
	printCurrency("BTC", "$", exchange.Btc.Satis, exchange.Btc.Alis)
}

func getSterlin() {
	exchange := getExchange()
	printCurrency("STERLIN", "₺", exchange.Gbp.Satis, exchange.Gbp.Alis)
}

func getEth() {
	exchange := getExchange()
	printCurrency("ETH", "$", exchange.Eth.Satis, exchange.Eth.Alis)
}
func getGold() {
	exchange := getExchange()
	printCurrency("GOLD", "₺", exchange.Gold.Satis, exchange.Gold.Alis)
}

func getExchange() *model.Exchange {
	exchange, err := exchange.GetExchange()
	if err != nil {
		fmt.Println(err)
	}
	return exchange
}

func printCurrency(currencyName string, currencyUnit string, buying string, sales string) {
	current_time := time.Now()

	fmt.Printf("[%s] [%s] : Buying: %s %s, Sales: %s₺ %s", currencyName, current_time.Format("2006-01-02 15:04:05"), buying, currencyUnit, sales, currencyUnit)

	fmt.Println()
}

func init() {
	rootCmd.AddCommand(get)
	get.Flags().BoolP("usd", "u", false, "Get USD Currency")
	get.Flags().BoolP("euro", "e", false, "Get Euro Currency")
	get.Flags().BoolP("sterlin", "s", false, "Get Sterlin Currency")
	get.Flags().BoolP("btc", "b", false, "Get BTC Currency")
	get.Flags().BoolP("eth", "", false, "Get ETH Currency")
	get.Flags().BoolP("gold", "g", false, "Get GOLD Currency (Gram)")

}
