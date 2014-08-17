# Example

Sample application for fluent-logger-go using [gin](https://github.com/gin-gonic/gin).

## Setup

```
$ go get github.com/gin-gonic/gin
$ gem install fluentd
$ gem install foreman
```

## Launch

```
$ foreman start
```

If you visit http://localhost:3000/foo, { "id": "foo" } is logged to fluentd.
