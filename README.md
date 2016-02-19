[![GoDoc](https://godoc.org/github.com/thomaso-mirodin/go-shorten?status.svg)](http://godoc.org/github.com/thomaso-mirodin/go-shorten)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/shawty/master/LICENSE)

## go-shorten: URL Shortener Service

This service stores and serves URL redirects

### Can I use it in production?

I dunno? I will be at some point.

### Why?

By itself, URL shortening is stupid useful.

But this project's parent <https://github.com/didip/shawty> existed to demonstrate:

* How concise [Go](http://golang.org/) is. [cloc](http://cloc.sourceforge.net/) shows that it could do URL shortening in only 125 lines.

* How slim Go is: 3MB RAM.

* How comprehensive Go standard library is.

* How easy it is to get up and running in Go. It took them about 1 hour from start to finish. Writing this README file took more time.

* How performant Go is:
    ```
    # Command  : ab -n 100000 -c 200 -k http://localhost:8080/dec/1
    # Processor: 2.26 GHz Intel Core 2 Duo  <-- Crummy 6 years old laptop

    Concurrency Level:      200
    Time taken for tests:   8.610 seconds
    Complete requests:      100000
    Failed requests:        0
    Non-2xx responses:      100000
    Keep-Alive requests:    100000
    Total transferred:      22400000 bytes
    HTML transferred:       7600000 bytes
    Requests per second:    11614.80 [#/sec] (mean)
    Time per request:       17.219 [ms] (mean)
    Time per request:       0.086 [ms] (mean, across all concurrent requests)
    Transfer rate:          2540.74 [Kbytes/sec] received
    ```
