package openweathermap

// Coord holds the longitude and latitude for the requested location
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Weather contains data related to general weather conditions
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main holds the current temperature, atmospehric pressure, humidity, and
// minimum and maximum temperature ranges
type Main struct {
	Temp      float64 `json:"temp"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	GrndLevel float64 `json:"grnd_level"`
}

// Wind specifies the wind direction and speed
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Clouds holds data related to cloudiness
type Clouds struct {
	All int `json:"all"`
}

// Sys holds general info about a request
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// CurrentWeatherData describes the current weather information for a location
type CurrentWeatherData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}
