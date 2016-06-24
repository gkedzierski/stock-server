stock-server
============

A socket based server polling Yahoo Finance API for stock market data.  
Polls data every second and immediately writes to the stream.  
Exposes `GET /quote/{TICKER}` endpoint.

## Configuration

Install normally in your GOPATH.

### Build

```
go build
```

### Test

```
go test
```

## TODO

 * Only write updated data to the stream
