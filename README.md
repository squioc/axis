# Axis

A golang library to manipulate generic dimensions (like time or distance).

Inspired by [timetools](https://github.com/mailgun/timetools) from [mailgun](https://github.com/mailgun).

## Usage:

```go
package main

import (
    "github.com/squioc/axis"
    "fmt"
    "time"
    "os"
)

func main() {
    provider := axis.NewFakeTime(2000)
    fmt.Println(provider.Current()) // 2000

    provider.Sleep(400) //Simulate a sleep
    fmt.Println(provider.Current()) // 2400

    C := make(chan int, 1)
    ch := provider.After(300)
    go func() {
        select {
            case position := <-ch:
                fmt.Println(position) // 2800
                C <- 1
            case <- time.After(time.Second):
                fmt.Fprintln(os.Stderr, "Timeout")
                C <- 1
        }
    } ()

    provider.Update(2800)
    <- C
}
```
