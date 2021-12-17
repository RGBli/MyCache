# MyCache
A LRU cache implemented in Golang.


</br>


### Usage
```go
import (
    "fmt"
    "github.com/RGBli/MyCache"
    "github.com/RGBli/MyCache/types"
)

func main() {
    // initialize a cache instance
    cache := MyCache.New(1 * 1024 * 1024)
    
    // put value in cache
    key := "lbw"
    value := types.NewString("23")
    cache.Set(key, value)
    
    // get value from cache
    v, ok := cache.GetString(key)
    fmt.Println(v, ok)
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
|value types|implement by yourself|5 common implemented types|

</br>

One thing to notice. RWLock and map are applied to mycache, so mycache works not as fast as map in Golang, but the gap is not that clear. You can run benchmark_test.go to get more infomation about it.
