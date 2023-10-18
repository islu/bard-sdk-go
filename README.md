# Google Bard <img src="https://www.gstatic.com/lamda/images/favicon_v1_150160cddff7f294ce30.svg" width="20px" /> SDK for Go

## Based on these GitHub Repository

- https://github.com/dsdanielpark/Bard-API (python)
- https://github.com/Allan-Nava/go-bard (go)
- https://github.com/ganeshk312/bard-go (go)

## Installation

```
go get github.com/islu/bard-sdk-go
```

## Authentication

1. Go to https://bard.google.com
2. F12 for console
3. Session: Application → Cookies → Copy the value of `__Secure-1PSID` cookie.

## Usage

Simple Usage
```go
package main

import "github.com/islu/bard-sdk-go/bard"

func main() {
    bot, err := bard.NewChatbot("BARD_API_KEY")
    if err != nil {
        // log.Fatalln(err)
        return
    }
    response := bot.Ask("your prompt")
}
```
