[![GoDoc](https://godoc.org/github.com/didip/shawty?status.svg)](http://godoc.org/github.com/didip/shawty)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/shawty/master/LICENSE)

## Shawty: URL Shortener Service

This service encodes URL in base-36 and store them in filesystem.

It has 3 features: shorten, unshorten, and redirect.


### Can I use it in production?

You need to implement a storage that can scale beyond one application server.


### Why?

By itself, URL shortening is quite useful.

But this project exists to demonstrate:

* How concise [Go](http://golang.org/) is. [cloc](http://cloc.sourceforge.net/) shows that this project contains only 125 lines.

* How slim Go is: 3MB RAM.

* How comprehensive Go standard library is.

* How easy it is to get up and running in Go. It took me about 1 hour from start to finish. Writing this README file took longer time.

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


### My other Go libraries

* [Tollbooth](https://github.com/didip/tollbooth): Simple middleware to rate-limit HTTP requests.

* [Gomet](https://github.com/didip/gomet): Simple HTTP client & server long poll library for Go. Useful for receiving live updates without needing Websocket.

* [Stopwatch](https://github.com/didip/stopwatch): A small library to measure latency of things. Useful if you want to report latency data to Graphite.

* [LaborUnion](https://github.com/didip/laborunion): A dynamic worker pool library.
