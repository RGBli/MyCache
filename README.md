# MyCache
---
A simple LRU cache implemented in Golang.


</br>


### Usage
```go
import (
    "github/RGBli/MyCache"
)

func main() {
    // define a type that can be cached
    type myValue struct {
    }
    
    // myValue implements Valuer interface so it can be cached
    func (v *myValue) Size() uint64 {
        return 1
    }
    
    // initialize a cache
    cache := MyCache.New(1 * 1024 * 1024)
    
    // put value in the cache
    value := myValue{}
    cache.Set("lbw", value)
    
    // get value from the cache
    v, ok := cache.Get("key")
}
```