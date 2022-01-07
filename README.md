# Go Rate Limiter

Description
This is an demo implementation of rate limiting 
in go using token bucket algorithm.

## Directory Structure :
1. bucket     -> Contains the core implementation of the token
                bucket algroithm
2. middleware -> Contains the implementation of rate limiting as middlware in
                API routes.
3. rules      ->  rules contains the interaction layer between rate limiter and corresponding
                  token buckets.

## How to run this project
```
$ go run .
```
