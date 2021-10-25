# go-minimal-cache
An in-memory key-value store library for Go, appropriate for single-machine applications. It provides thread-safe `map[string]interface{}` and allow fast acces.

Any object can be stored and the cache can be safely used by multiple goroutines.

Although go-cache isn't meant to be used as a persistent datastore, the entire
cache can be saved to and loaded from a file. The existing values of the keys on cache 
are saved to the file in json format with every 60 second intervals. Time duration to save 
file and file path can be changed by the developers with methods `SetPath` and `SetFrequency`.

Also available to see an image of the memory with `GetValue`.

### Installation

`go get github.com/HarunBuyuktepe/go-minimal-cache`

### Usage

```go
import (
    "fmt"
    "github.com/HarunBuyuktepe/go-minimal-cache"
)

func main() {
    // Create a cache 
    c := cache.New()

    // Set the value of the key "foo" to "bar"
    c.Set("foo", "bar")

    // Set the value of the key "foo" to 1
    c.Set("foo", 1)

    // Get the string associated with the key "foo" from the cache
    foo := c.Get("foo")
    fmt.Println(foo)

    // Delete the string associated with the key "foo" from the cache
    c.Delete("foo")
    // It will print "" only
    fmt.Println(foo)

    // DeleteAll provides to clean cache
    c.DeleteAll()

    // Set frequency 120 seconds to write cache to json file
    c.SetFrequency(120)

    // Get frequency is available
    fmt.Println(c.GetFrequency())

    // Set path test.json
    c.SetPath("test.json")

    // Get path is also available
    fmt.Println(c.GetPath())

    // GetValue send a picture of the cache and return key value relation with json formatted 
    fmt.Println(c.GetValue())

  }
}
```

### Example

May you see the sample implementation? -> [click me](https://github.com/HarunBuyuktepe/go-restapi-cache)

