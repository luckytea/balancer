# balancer

Lazy round-robin balancer.

## Example

```golang
var hosts = []string("host1","host2")

b := balancer.New(hosts)

url := b.GetHost()

http.Get(url)
```
