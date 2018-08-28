package main

import (
    "encoding/json"
    "fmt"
    "os"
    "owm/fiveday"
)

func printFiveDayData(forecastData fiveday.Data) {
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(forecastData)
}

func main() {
    filename := os.Args[1]
    var forecastData fiveday.Data
    forecastData.FromFile(filename)

    if json, err := forecastData.ToJson(); err == nil {
        fmt.Println(json)
    } else {
        fmt.Println("error:", err)
    }
}
