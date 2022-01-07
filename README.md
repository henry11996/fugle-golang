# Fugle-golang

[![Go Report Card](https://goreportcard.com/badge/github.com/henry11996/fugle-golang)](https://goreportcard.com/report/github.com/henry11996/fugle-golang)
[![GitHub release](https://img.shields.io/github/v/tag/henry11996/fugle-golang.svg?label=release)](https://github.com/henry11996/fugle-golang/releases)

## What is fugle golang?
Is a [Fugle Api](https://developer.fugle.tw/) library.

## Installation
Install
```
go get github.com/henry11996/fugle-golang
```
Updates
```
go get -u -v github.com/henry11996/fugle-golang/...
```
## Quick start
```go
package main

import (
	"log"
	"github.com/henry11996/fugle-golang/fugle"
)

fugleToken := "your_token"

clientOption := fugle.ClientOption{
    ApiToken: fugleToken,
}

client, err := fugle.NewFugleClient(clientOption)

if err != nil {
    log.Panic(err)
}

//IX0001 台灣加權指數
meta, _ := client.Meta("IX0001", false)

log.Print(meta)
```

## Feature
* [x] HTTP API
* [ ] WebSocket
* [ ] Unit Tests