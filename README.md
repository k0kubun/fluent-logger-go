# fluent-logger-go

Goroutine based asynchronous event logger for Fluentd

## Features

* Channel based non-blocking logging interface
* Queued events are periodically sent to fluentd altogether
* Single goroutine owns a role of logging, thus no risk for duplicated logging

## Installation

```
$ go get github.com/k0kubun/fluent-logger-go
```

## Usage

```
import "github.com/k0kubun/fluent-logger-go"
```

pending

### Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/fluent-logger-go

## Related projects

You would like to use another one for some use case.

* [t-k/fluent-logger-golang](https://github.com/t-k/fluent-logger-golang)
