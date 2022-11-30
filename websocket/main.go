package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"os"

	"golang.org/x/net/websocket"
)

/// How to connect to Bitflyer JSON-RPC api

type RPCParam struct {
	Channel string `json:"channel"`
}

type RPCRequest struct {
	ID     string   `json:"id"`
	Method string   `json:"method"`
	Params RPCParam `json:"params"`
}

type ErrorObj struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RPCResponse struct {
	ID     string   `json:"id"`
	Result bool     `json:"result"`
	Error  ErrorObj `json:"error"`
}

func NewJsonRPC(method string, channel string) *RPCRequest {
	return &RPCRequest{
		ID:     fmt.Sprint(rand.Int()),
		Method: method,
		Params: RPCParam{Channel: channel},
	}
}

func (j *RPCRequest) MarshallJSON() ([]byte, error) {
	return json.Marshal(struct {
		jsonrpc string
		*RPCRequest
	}{
		jsonrpc:    "2.0",
		RPCRequest: j,
	})
}

func connectWS() *websocket.Conn {
	log.Println("start")
	origin := "http://localhost"
	url := "wss://ws.lightstream.bitflyer.com/json-rpc"

	log.Println("connecting")
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")

	return ws
}

type OrderBookSnapshot struct {
	Jsonrpc string `json:"_"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Message struct {
			Asks []struct {
				Price float64 `json:"price"`
				Size  float64 `json:"size"`
			} `json:"asks"`
			Bids []struct {
				Price float64 `json:"price"`
				Size  float64 `json:"size"`
			} `json:"bids"`
			MidPrice float64 `json:"mid_price"`
		} `json:"message"`
	} `json:"params"`
}

type Ticker struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Message struct {
			BestAsk         float64   `json:"best_ask"`
			BestAskSize     float64   `json:"best_ask_size"`
			BestBid         float64   `json:"best_bid"`
			BestBidSize     float64   `json:"best_bid_size"`
			Ltp             float64   `json:"ltp"`
			MarketAskSize   float64   `json:"market_ask_size"`
			MarketBidSize   float64   `json:"market_bid_size"`
			ProductCode     string    `json:"product_code"`
			State           string    `json:"state"`
			TickID          float64   `json:"tick_id"`
			Timestamp       time.Time `json:"timestamp"`
			TotalAskDepth   float64   `json:"total_ask_depth"`
			TotalBidDepth   float64   `json:"total_bid_depth"`
			Volume          float64   `json:"volume"`
			VolumeByProduct float64   `json:"volume_by_product"`
		} `json:"message"`
	} `json:"params"`
}

type Execution struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Message struct {
			ID                         int       `json:"id"`
			Side                       string    `json:"side"`
			Price                      int       `json:"price"`
			Size                       float64   `json:"size"`
			ExecDate                   time.Time `json:"exec_date"`
			BuyChildOrderAcceptanceID  string    `json:"buy_child_order_acceptance_id"`
			SellChildOrderAcceptanceID string    `json:"sell_child_order_acceptance_id"`
		} `json:"message"`
	} `json:"params"`
}

const (
	channelSnapshot = "lightning_board_snapshot_%s"
	channelTicker   = "lightning_ticker_%s"
	channelExec     = "lightning_execution_%s"
)

func snapshot(s Symbol) string {
	return fmt.Sprintf(channelSnapshot, s)
}

func ticker(s Symbol) string {
	return fmt.Sprintf(channelTicker, s)
}

func exec(s Symbol) string {
	return fmt.Sprintf(channelExec, s)
}

type Symbol string

const (
	btcfx Symbol = "FX_BTC_JPY"
)

func main() {
	ws := connectWS()
	channel := ticker(btcfx)
	// channel := exec(btcfx)

	if err := websocket.JSON.Send(ws, NewJsonRPC("subscribe", channel)); err != nil {
		log.Fatal(err)
		return
	}

	var v RPCResponse
	if err := websocket.JSON.Receive(ws, &v); err != nil {
		log.Fatal(err)
		return
	}
	if !v.Result {
		log.Println(v)
		log.Fatal("failed to subscribe")
		return
	}
	log.Println("subscribed")

	defer func(ws *websocket.Conn) {
		if err := websocket.JSON.Send(ws, NewJsonRPC("unsubscribe", channel)); err != nil {
			log.Fatal(err)
		}
		var v RPCResponse
		if err := websocket.JSON.Receive(ws, &v); err != nil {
			log.Fatal(err)
			return
		}

		if !v.Result {
			log.Fatal("failed to unsubscribe")
			return
		}
	}(ws)

	errChan := make(chan error)
	c := make(chan Ticker)
	rec := func(ws *websocket.Conn) {
		log.Println("getting")
		var v Ticker
		if err := websocket.JSON.Receive(ws, &v); err != nil {
			errChan <- err
			return
		}
		c <- v
	}

	go rec(ws)

	t1 := time.Now()
	s := make([]Ticker, 0)
	for {
		var v Ticker
		if err := websocket.JSON.Receive(ws, &v); err != nil {
			errChan <- err
			return
		}
		s = append(s, v)
    fmt.Println(len(s))
		if len(s) == 500 {
			file, _ := os.Create("test_ticker.json")
			d, _ := json.Marshal(s)
			fmt.Fprint(file, string(d))
			fmt.Println(time.Now().Sub(t1).Seconds())
			break
		}

		// select {
		// case err := <-errChan:
		// 	log.Fatal(err)
		// 	return
		// case v := <-c:
		// 	// log.Println("received response")
		// 	// file, _ := os.Create("test_ticker.json")
		// 	// d, _ := json.Marshal(v)
		// 	// fmt.Fprint(file, string(d))
		// 	s = append(s, v)
		// 	if len(s) < 500 {
		// 		go rec(ws)
		// 	} else {
		// 		file, _ := os.Create("test_ticker.json")
		// 		d, _ := json.Marshal(s)
		// 		fmt.Fprint(file, string(d))
		// 		fmt.Println(time.Now().Sub(t1).Seconds())
		// 		return
		// 	}
		// }
	}
}
