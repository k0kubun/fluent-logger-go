# fluent-logger-go

Goroutine based asynchronous event logger for Fluentd

## Features

* Channel based non-blocking logging interface
* Queued events are periodically sent to fluentd altogether

## Installation

```
$ go get github.com/k0kubun/fluent-logger-go
```

## Usage

```go
package main

import "github.com/gin-gonic/gin"
import "github.com/k0kubun/fluent-logger-go"

var logger *fluent.Logger

func logId(c *gin.Context) {
	id := c.Params.ByName("id")
	logger.Post("idlog", map[string]string{"id": id})
	c.String(200, "Logged id")
}

func main() {
	logger = fluent.NewLogger(fluent.Config{})
	r := gin.Default()
	r.GET("/:id", logId)
	r.Run(":3000")
}
```

### Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/fluent-logger-go

## Related projects

You would like to use another one for some use case.

* [t-k/fluent-logger-golang](https://github.com/t-k/fluent-logger-golang)
