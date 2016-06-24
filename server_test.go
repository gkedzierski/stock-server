// stock-server
// Socket based server streaming stock data for a given symbol.
// Used by frontend: https://github.com/gkedzierski/stock-viewer
//
// Copyright (c) 2016 Greg Kedzierski
// http://gregkedzierski.com
// greg@gregkedzierski.com

package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "net/http/httptest"
  "testing"
)

var m *mux.Router
var w *httptest.ResponseRecorder

func setup() {
  m = mux.NewRouter()
  addEndpointHandlers(m)

  w = httptest.NewRecorder()
}

func TestInvalidEndpoint(t *testing.T) {
  setup()

  req, err := http.NewRequest("GET", "/nonexistent/endpoint", nil)
  if err != nil {
    t.Fatal("Creating 'GET /nonexistent/endpoint' request failed.")
  }
  m.ServeHTTP(w, req)

  if w.Code != http.StatusNotFound {
    t.Errorf("/nonexistent/endpoint endpoint didn't return %v", http.StatusNotFound)
  }
}
