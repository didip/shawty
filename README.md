## URL Shortener Service

This simple service Base36 encodes URL and store them in file system.

It has 3 features: shorten, unshorten, and redirect.


### Dependency

`github.com/gorilla/mux`

### Why?

By itself, URL shortening is quite useful.

But this project exists to demonstrate:

* How short [Go](http://golang.org/) is: 80 lines.

* How slim Go is: 3MB RAM.

* How performant Go is:
    ```
    # Command  : ab -n 100000 -c 200 -k http://localhost:8080/dec/4
    # Processor: 2.26 GHz Intel Core 2 Duo

    Concurrency Level:      200
    Time taken for tests:   9.678 seconds
    Complete requests:      100000
    Failed requests:        0
    Write errors:           0
    Keep-Alive requests:    100000
    Total transferred:      16000000 bytes
    HTML transferred:       1900000 bytes
    Requests per second:    10333.08 [#/sec] (mean)
    Time per request:       19.355 [ms] (mean)
    Time per request:       0.097 [ms] (mean, across all concurrent requests)
    Transfer rate:          1614.54 [Kbytes/sec] received
    ```