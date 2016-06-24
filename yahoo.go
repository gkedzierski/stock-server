// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

// structure defining JSON response from Yahoo Finance API
type YahooQuoteResponse struct {
  List YahooQuoteResponseList `json:"list"`
}

type YahooQuoteResponseList struct {
  Meta YahooQuoteResponseMeta `json:"meta"`
  Resources []YahooQuoteResourceContainer `json:"resources"`
}

type YahooQuoteResponseMeta struct {
  Type string `json:"type"`
  Start int32 `json:"start"`
  Count int32 `json:"count"`
}

type YahooQuoteResourceContainer struct {
  Resource YahooQuoteResource `json:"resource"`
}

type YahooQuoteResource struct {
    ClassName string `json:"classname"`
    Fields YahooQuoteResourceFields `json:"fields"`
}

type YahooQuoteResourceFields struct {
    Change string `json:"change"`
    ChangePercentage string `json:"chg_percent"`
    DayHigh string `json:"day_high"`
    DayLow string `json:"day_low"`
    IssuerName string `json:"issuer_name"`
    IssuerNameLang string `json:"issuer_name_lang"`
    Name string `json:"name"`
    Price string `json:"price"`
    Symbol string `json:"symbol"`
    Timestamp string `json:"ts"`
    Type string `json:"type"`
    UTCTime string `json:"utctime"`
    Volume string `json:"volume"`
    YearHigh string `json:"year_high"`
    YearLow string `json:"year_low"`
}
