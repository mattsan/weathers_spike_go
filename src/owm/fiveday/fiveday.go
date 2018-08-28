package fiveday

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
