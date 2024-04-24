package main

import "dca-bitcoin/services"

func main() {
	services.BuyBitcoin()
	go services.StartCronJob()
	select {}
}
