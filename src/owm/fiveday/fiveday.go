package fiveday

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
)

type Volume map[string]float64

type Wind struct {
    Speed float64
    Deg float64
}

type Clouds struct {
    All int
}

type Weather struct {
      Id int
      Main string
      Description string
      Icon string
}

type Main struct {
    Temp float64
    Temp_min float64
    Temp_max float64
    Pressure float64
    Sea_level float64
    Grnd_level float64
    Humidity int
    Temp_kf float64
}

type Item struct {
    Dt int
    Main Main
    Weather []Weather
    Clouds Clouds
    Wind Wind
    Rain Volume
    Snow Volume
    Dt_txt string
}

type Coord struct {
    Lat float64
    Lan float64
}

type City struct {
    Id int
    Name string
    Coord Coord
    Country string
}

type Data struct {
    Cod string
    Message float64
    City City
    Cnt int
    List []Item
}

func (fiveDayData *Data) FromReader(reader io.Reader) {
    decoder := json.NewDecoder(reader)

    err := decoder.Decode(fiveDayData)

    if err != nil {
        fmt.Println("DECODE ERROR:", err)
    }
}

func (fiveDayData *Data) FromFile(filename string) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("FILE READ ERROR:", err)
    }

    fiveDayData.FromReader(io.Reader(bytes.NewBuffer(content)))
}

func (fiveDayData *Data) ToJson() string {
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
