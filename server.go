// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main;

import (
  "github.com/gorilla/websocket"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "log"
  "time"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  // allow all origins
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

// starts a server on port 3000
func startServer() {
  rtr := mux.NewRouter()
  addEndpointHandlers(rtr)
  http.Handle("/", rtr);

  log.Println("Listening on port 3000...")
  http.ListenAndServe(":3000", nil)
}

func addEndpointHandlers(rtr *mux.Router) {
  rtr.HandleFunc("/quote/{symbol:[a-zA-Z0-9]+}", handleQuoteRequest).Methods("GET")
}

// handles incoming request
func handleQuoteRequest(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)

  // retrieve params from address
  params := mux.Vars(r)
  symbol := params["symbol"]

  if err != nil {
    log.Println(err)
    return
  }
  log.Printf("Client subscribed from %s asking for %s\n", r.RemoteAddr, symbol)

  // infinite loop polling stock API every second
  for {
    // fetch data from API
    quote, err := fetchQuote(symbol)
    if err != nil {
      log.Println(err)
      break;
    }

    // encode object to JSON
    jsonOutput, err := json.Marshal(quote)
    if err != nil {
      log.Println(err)
      return
    }

    // send encoded JSON back to client
    err = conn.WriteMessage(websocket.TextMessage, jsonOutput)
    if err != nil {
      log.Println(err)
      break
    }

    // wait a second before next poll
    time.Sleep(time.Second)
  }

  log.Printf("Client asking for %s unsubscribed\n", symbol)
}
