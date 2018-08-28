package main

import (
    "encoding/json"
    "fmt"
    "os"
    "owm/fiveday"
)

func printFiveDayData(fiveDayData fiveday.Data) {
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(fiveDayData)
}

func main() {
    filename := os.Args[1]
    var fiveDayData fiveday.Data
    fiveDayData.FromFile(filename)

    fmt.Println(fiveDayData.ToJson())
}
