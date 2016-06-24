// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

import (
  "testing"
)

func TestQuoteDataFetch(t *testing.T) {
  dataObject, err := fetchQuote("MSFT")
  if err != nil {
    t.Errorf("Fetching quote data failed: %v", err)
  }

  if dataObject.Name != "Microsoft Corporation" {
    t.Error("Correct data for MSFT could not be fetched.")
  }
}
