// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// fetches historical data for a given ticker symbol from
// local CSV file
func fetchHistoricalData(symbol string) ([]CandleStick, error) {
	var historicalData []CandleStick

	// open CSV file
	f, err := os.Open("data/" + strings.ToUpper(symbol) + ".csv")
	if err != nil {
		return nil, err
	}

	// read
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// parse timestamp
		form := "2006-01-02"
		timeObject, err := time.Parse(form, record[0])
		if err != nil {
			continue
		}
		timestamp := timeObject.Unix()

		// parse open value
		open, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			continue
		}

		// parse close value
		close, err := strconv.ParseFloat(record[4], 32)
		if err != nil {
			continue
		}

		// parse high value
		high, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			continue
		}

		// pass low value
		low, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			continue
		}

		// append datapoint to slice
		historicalData = append(historicalData, CandleStick{
			Timestamp: timestamp,
			Open:      open,
			Close:     close,
			High:      high,
			Low:       low,
		})
	}

	return historicalData, nil
}
