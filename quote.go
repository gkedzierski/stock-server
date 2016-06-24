// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

import (
  "net/http"
  "encoding/json"
)

// fetches quote using Yahoo Finance API and returns hydrated
// structure
func fetchQuote(symbol string) (*YahooQuoteResourceFields, error) {
  // make a HTTP request
  res, err := http.Get("http://finance.yahoo.com/webservice/v1/symbols/" + symbol + "/quote?format=json&view=detail")
  if err != nil {
    return nil, err
  }

  // decode JSON - hydrate structure
  var decodedResponse = new(YahooQuoteResponse)
  err = json.NewDecoder(res.Body).Decode(decodedResponse)
  if err != nil {
    return nil, err
  }

  return &decodedResponse.List.Resources[0].Resource.Fields, nil
}
