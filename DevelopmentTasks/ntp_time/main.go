package main

import (
    "fmt"
    "log"
    "os"
    "github.com/beevik/ntp"
)

func main() {
    time, err := ntp.Time("pool.ntp.org")
    if err != nil {
        log.Printf("Error fetching time: %v", err)
        os.Exit(1)
    }
    fmt.Printf("Current time: %v\n", time)
}
