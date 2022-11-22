package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	// "os"

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
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Message struct {
			Asks []struct {
				Price int     `json:"price"`
				Size  float64 `json:"size"`
			} `json:"asks"`
			Bids []struct {
				Price int     `json:"price"`
				Size  float64 `json:"size"`
			} `json:"bids"`
			MidPrice float64 `json:"mid_price"`
		} `json:"message"`
	} `json:"params"`
}

func main() {
	ws := connectWS()

	if err := websocket.JSON.Send(ws, NewJsonRPC("subscribe", "lightning_board_snapshot_FX_BTC_JPY")); err != nil {
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
		if err := websocket.JSON.Send(ws, NewJsonRPC("unsubscribe", "lightning_board_snapshot_FX_BTC_JPY")); err != nil {
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
	c := make(chan OrderBookSnapshot)
	rec := func(ws *websocket.Conn) {
		log.Println("getting order book")
		var v OrderBookSnapshot
		if err := websocket.JSON.Receive(ws, &v); err != nil {
			errChan <- err
			return
		}
		// file, _ := os.Create("test.json")
		// d, _ := json.Marshal(v)
		// fmt.Fprint(file, string(d))
		c <- v
	}

	go rec(ws)

	for {
		select {
		case err := <-errChan:
			log.Fatal(err)
			return
		case v := <-c:
			log.Println("received response")
			log.Println(v)
			go rec(ws)
		}
	}
}
