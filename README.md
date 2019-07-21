# OpenWeatherMap Go Wrapper

Package `openweathermap` implements a client for the [OpenWeatherMap API](https://openweathermap.org/api).

Note: At this time, this package supports the [Current weather data](https://openweathermap.org/current) endpoint only.

Before using the OpenWeatherMap API, you need to [register here](http://home.openweathermap.org/users/sign_up) and read the [API Guidelines](https://openweathermap.org/current).

## Installation

```bash
go get github.com/ayoisaiah/openweathermap
```

Then import the package into your project:

```go
import (
    github.com/ayoisaiah/openweathermap
)
```

## Creating an instance

An API Key from OpenWeatherMap is required

```go
owm, err := openweathermap.New("YOUR_API_KEY")
if err != nil {
  log.Fatalln(err)
}

// do something with `owm`
```

You may optionally provide a custom HTTP client:

```go
owm, err := openweathermap.New("{YOUR_API_KEY}", WithHTTPClient(http.DefaultClient))
if err != nil {
  log.Fatalln(err)
}

// do something with `owm`
```

Or set your preferred temperature unit (one of `F` (Fahrenheit), `C` (Celsuis), `K`(Kelvin)). The default is Kelvin.

```go
owm, err := openweathermap.New("{YOUR_API_KEY}", WithUnit("C"))
if err != nil {
  log.Fatalln(err)
}

// do something with `owm`
```

Or change the default language (English), to any of the [supported languages](https://github.com/ayoisaiah/openweathermap#supported-languages) below:

```go
owm, err := openweathermap.New("{YOUR_API_KEY}", WithLang("fr"))
if err != nil {
  log.Fatalln(err)
}

// do something with `owm`
```

## Get the current weather forecast

### By city name

```go
forecast, err := owm.GetCurrentByCityName("Lagos")
if err != nil {
  fmt.Println(err)
  return
}

fmt.Printf("%+v", forecast)
```

### By coordinates

```go
forecast, err := owm.GetCurrentByCoords(&openweathermap.Coord{
  Lat: 8.5,
  Lon: 4.5,
})
if err != nil {
  fmt.Println(err)
  return
}

fmt.Printf("%+v", forecast)
```

### By city ID

```go
forecast, err := owm.GetCurrentByID(4513583)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Printf("%+v", forecast)
```

### By Zip code

```go
forecast, err := owm.GetCurrentByZipCode(94040, "us")
if err != nil {
  fmt.Println(err)
  return
}

fmt.Printf("%+v", forecast)
```

## Supported languages

English - en, Russian - ru, Italian - it, Spanish - es (or sp), Ukrainian - uk (or ua), German - de, Portuguese - pt, Romanian - ro, Polish - pl, Finnish - fi, Dutch - nl, French - fr, Bulgarian - bg, Swedish - sv (or se), Chinese Traditional - zh_tw, Chinese Simplified - zh (or zh_cn), Turkish - tr, Croatian - hr, Catalan - ca

## Testing

Ensure the `OPENWEATHERMAP_KEY` environmental variable is set before running the tests

```bash
env OPENWEATHERMAP_KEY="YOUR_API_KEY" go test
```

## Licence

MIT
