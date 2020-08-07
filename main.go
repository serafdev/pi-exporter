package main

import (
        "log"
        "fmt"
        "github.com/d2r2/go-dht"
)

func main() {

        temperature, humidity, retried, err :=
                dht.ReadDHTxxWithRetry(dht.DHT11, 3, false, 10)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
                temperature, humidity, retried)
}
