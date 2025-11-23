### Problem: API Rate Limiter for Concurrent HTTP Requests
You are asked to implement a function that makes concurrent HTTP GET requests to multiple URLs but respects a **concurrency limit**.
#### Requirements:
* The function signature should look like this (you can adapt slightly if needed):
```go
func FetchURLs(urls []string, limit int) []string
```
* **Parameters:**
* `urls`: a list of URLs to send GET requests to.
* `limit`: the maximum number of concurrent HTTP requests allowed at any time.

* **Behavior:**
1. The function should make HTTP GET requests to all given URLs.
2. At most `limit` requests can be in-flight concurrently.
3. The function should collect the response bodies (or errors) and return them in the same order as the input URLs.

4. For simplicity purposes the request sending part might be mocked/simulated.