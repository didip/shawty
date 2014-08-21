## URL Shortener Service

This service encodes URL in base-36 and store them in filesystem.

It has 3 features: shorten, unshorten, and redirect.


### Can I use it in production?

You may want to handle the errors better before using it in production.

To scale out beyond 1 server, you can put the data in NFS/Ceph/Gluster.


### Why?

By itself, URL shortening is quite useful.

But this project exists to demonstrate:

* How short [Go](http://golang.org/) is: 80 lines.

* How slim Go is: 3MB RAM.

* How comprehensive Go standard library is. If `net/http` supports basic pattern matching, this project would be 100% based on standard library.

* How quick it is to get up and running in Go. It took me about 1 hour from start to finish. Writing this README file took longer time.

* How performant Go is:
    ```
    # Command  : ab -n 100000 -c 200 -k http://localhost:8080/dec/4
    # Processor: 2.26 GHz Intel Core 2 Duo  <-- Crummy 6 years old laptop

    Concurrency Level:      200
    Time taken for tests:   8.512 seconds
    Complete requests:      100000
    Failed requests:        0
    Write errors:           0
    Keep-Alive requests:    100000
    Total transferred:      16005920 bytes
    HTML transferred:       1900703 bytes
    Requests per second:    11747.64 [#/sec] (mean)
    Time per request:       17.025 [ms] (mean)
    Time per request:       0.085 [ms] (mean, across all concurrent requests)
    Transfer rate:          1836.25 [Kbytes/sec] received
    ```