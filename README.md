# MyCache
A thread-safe LRU cache implemented in Golang.


</br>


### Usage
```go
import (
    "fmt"
    "github.com/RGBli/MyCache"
    "github.com/RGBli/MyCache/skiplist"
)

func main() {
    // initialize a cache instance
    cache := MyCache.New(1 * 1024 * 1024)

    // use database "test"
    db := cache.Use("test")
    
    // put value in cache
    key := "lbw"
    value := types.NewString("23")
    db.Set(key, value)
    
    // get value from cache
    v, ok := db.GetString(key)
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

One thing to notice. RWLock is applied to mycache, so mycache works not as fast as map in Golang, but the gap is not that clear. You can run benchmark_test.go to get more information about it.

</br>

### Todo
1)Data persistence