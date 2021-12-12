# MyCache
A LRU cache implemented in Golang.


</br>


### Usage
```go
import (
    "github.com/RGBli/MyCache"
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

</br>

### Features
|     |map|MyCache|
|:---:|:---:|:---:|
|thread-safe|no|yes|
|maximum-capacity|no|yes|
|scalable|no|yes|
|TTL|no|yes|

</br>

One thing to notice. RWLock and map are applied to mycache, so mycache works not as fast as map in Golang, but the gap is not that clear. You can run benchmark_test.go to get more infomation about it.
