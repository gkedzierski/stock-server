// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

// structure defining single datapoint for a candlestick chart
type CandleStick struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	Low       float64 `json:"low"`
	High      float64 `json:"high"`
}
