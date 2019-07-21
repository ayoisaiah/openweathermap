package openweathermap

import "net/http"

// OpenWeatherMap wraps the entire OpenWeatherMap API
type OpenWeatherMap struct {
	client *http.Client
	key    string
	unit   string
	lang   string
}

// Options represents OpenWeatherMap options
type Options func(owm *OpenWeatherMap) error

// WithUnit sets a custom unit when creating OpenWeatherMap instance
func WithUnit(unit string) Options {
	units := map[string]string{
		"C": "metric",
		"F": "imperial",
		"K": "standard",
	}

	return func(owm *OpenWeatherMap) error {
		if _, exists := units[unit]; !exists {
			return &IllegalArgumentError{ErrString: "Invalid unit"}
		}

		owm.unit = units[unit]
		return nil
	}
}

// WithLang sets custom language when creating OpenWeatherMap instance
func WithLang(lang string) Options {
	langCodes := []string{"en", "ru", "it", "es", "sp", "uk", "ua", "de", "pt",
		"ro", "pl", "fi", "nl", "fr", "bg", "sv", "se", "tr", "hr", "ca", "zh_tw",
		"zh", "zh_cn"}

	return func(owm *OpenWeatherMap) error {
		v := Contains(langCodes, lang)
		if !v {
			return &IllegalArgumentError{ErrString: "Invalid language code"}
		}

		owm.lang = lang
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client when creating OpenWeatherMap
// instance
func WithHTTPClient(client *http.Client) Options {
	return func(owm *OpenWeatherMap) error {
		if client == nil {
			return &IllegalArgumentError{ErrString: "Invalid HTTP Client"}
		}

		owm.client = client
		return nil
	}
}

// New returns a new OpenWeatherMap instance
func New(key string, opts ...Options) (*OpenWeatherMap, error) {
	if key == "" {
		return nil, &IllegalArgumentError{ErrString: "API Key cannot be empty"}
	}

	owm := &OpenWeatherMap{
		client: http.DefaultClient,
		key:    key,
		unit:   "metric",
		lang:   "en",
	}

	for _, opt := range opts {
		if opt == nil {
			return nil, &IllegalArgumentError{ErrString: "Options cannot be nil"}
		}

		err := opt(owm)
		if err != nil {
			return nil, err
		}
	}

	return owm, nil
}
