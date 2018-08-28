package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "owm/fiveday"
)

func printFiveDayData(fiveDayData fiveday.Data) {
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(fiveDayData)
}

func encodeToJson(fiveDayData fiveday.Data) string {
    var buffer bytes.Buffer
    encoder := json.NewEncoder(io.Writer(&buffer))
    err := encoder.Encode(fiveDayData)
    if err == nil {
        return buffer.String()
    } else {
        fmt.Println("ENCODE ERROR:", err)
        return "ERROR"
    }
}

func decodeFromJson(filename string) fiveday.Data {
    var fiveDayData fiveday.Data

    content, err := ioutil.ReadFile(filename)
    if err == nil {
        decoder := json.NewDecoder(io.Reader(bytes.NewBuffer(content)))

        err := decoder.Decode(&fiveDayData)
        if err == nil {
            return fiveDayData
        } else {
            fmt.Println("DECODE ERROR:", err)
        }
    } else {
        fmt.Println("FILE READ ERROR:", err)
    }
    return fiveDayData
}

func main() {
    filename := os.Args[1]

    fiveDayData := decodeFromJson(filename)

    // printFiveDayData(fiveDayData)
    fmt.Println(encodeToJson(fiveDayData))
}
