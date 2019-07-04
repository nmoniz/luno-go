package main

import (
	"log"
	"time"

	"github.com/luno/luno-go/streaming"
)

func printOrderBook(c *streaming.Conn) {
	seq, bids, asks := c.OrderBookSnapshot()
	seq, status := c.MarketStatus()
	log.Printf("%d: %s %v %v\n", seq, status, bids[0], asks[0])
}

func main() {
	keyID := "key here"
	keySecret := "key secret here"

	c, err := streaming.Dial(keyID, keySecret, "XBTZAR", streaming.WithConnectCallback(printOrderBook))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		time.Sleep(time.Minute)
		printOrderBook(c)
	}
}
