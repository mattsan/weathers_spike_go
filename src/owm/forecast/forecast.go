package forecast

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
)

type Volume struct {
    Per3hours float64 `json:"3h"`
}

type Wind struct {
    Speed float64 `json:"speed"`
    Deg float64 `json:"deg"`
}

type Clouds struct {
    All int `json:"all"`
}

type Weather struct {
      Id int `json:"id"`
      Main string `json:"main"`
      Description string `json:"description"`
      Icon string `json:"icon"`
}

type Main struct {
    Temp float64 `json:"temp"`
    TempMin float64 `json:"temp_min"`
    TempMax float64 `json:"temp_max"`
    Pressure float64 `json:"pressure"`
    SeaLevel float64 `json:"sea_level"`
    GrndLevel float64 `json:"grnd_level"`
    Humidity int `json:"humidity"`
    TempKf float64 `json:"temp_kf"`
}

type Item struct {
    Dt int `json:"dt"`
    Main Main `json:"main"`
    Weather []Weather `json:"weather"`
    Clouds Clouds `json:"clouds"`
    Wind Wind `json:"wind"`
    Rain Volume `json:"rain"`
    Snow Volume `json:"snow"`
    DtTxt string `json:"dt_txt"`
}

type Coord struct {
    Lat float64 `json:"lat"`
    Lan float64 `json:"lan"`
}

type City struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Coord Coord `json:"coord"`
    Country string `json:"Country"`
}

type Data struct {
    Cod string `json:"cod"`
    Message float64 `json:"message"`
    City City `json:"city"`
    Cnt int `json:"cnt"`
    List []Item `json:"list"`
}

func (data *Data) FromReader(reader io.Reader) error {
    decoder := json.NewDecoder(reader)

    return decoder.Decode(data)
}

func (data *Data) FromFile(filename string) error {
    if content, err := ioutil.ReadFile(filename); err == nil {
        return data.FromReader(io.Reader(bytes.NewBuffer(content)))
    } else {
        return err
    }
}

func (data *Data) ToJson() (string, error) {
    var buffer bytes.Buffer
    encoder := json.NewEncoder(io.Writer(&buffer))
    err := encoder.Encode(data)
    return buffer.String(), err
}
