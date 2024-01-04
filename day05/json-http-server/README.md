## Json Http Server
```
> curl -iX GET localhost:8080
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 20 Dec 2023 15:13:54 GMT
Content-Length: 19

404 page not found
```

```go
	curl -X POST localhost:8080 -d '{"activity": {"description": "christmas eve class", "time":"2021-12-24T12:42:31Z"}}'
```
```go
 curl -X GET localhost:8080 -d '{"id": 0}'                                                                           
{"activity":{"time":"2021-12-24T12:42:31Z","description":"christmas eve class","id":0}}

```

Ref:
- https://earthly.dev/blog/golang-http/
