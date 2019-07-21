package openweathermap

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// IllegalArgumentError occurs when the argument to a function are invalid
type IllegalArgumentError struct {
	ErrString string
}

func (e IllegalArgumentError) Error() string {
	return e.ErrString
}

// JSONDecodingError occurs when decoding response body to JSON fails
type JSONDecodingError struct {
	ErrString string
}

func (e JSONDecodingError) Error() string {
	return e.ErrString
}

// AuthorizationError occurs for an Unauthorized request
type AuthorizationError struct {
	ErrString string
}

func (e AuthorizationError) Error() string {
	return e.ErrString
}

// NotFoundError occurs when the resource queried returns a 404.
type NotFoundError struct {
	ErrString string
}

func (e NotFoundError) Error() string {
	return e.ErrString
}

// RateLimitError occurs when rate limit is reached for the API key.
type RateLimitError struct {
	ErrString string
}

func (e RateLimitError) Error() string {
	return e.ErrString
}

// CheckForErrors checks if a non success http status code is retured by an API
// request and returns the appropriate error
func CheckForErrors(resp *http.Response) error {
	buf, err := ioutil.ReadAll(resp.Body)

	// Because you can't read from an io.ReadCloser type twice unless you restore
	// the content
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200, 201, 202, 204, 205:
		return nil
	case 401:
		return &AuthorizationError{ErrString: errStringHelper(resp.StatusCode, "Unauthorized request", &buf)}
	case 403:
		return &AuthorizationError{ErrString: errStringHelper(resp.StatusCode, "Access forbidden request", &buf)}

	case 404:
		return &NotFoundError{ErrString: errStringHelper(resp.StatusCode, "The requested resource was not found", &buf)}
	case 429:
		return &RateLimitError{ErrString: errStringHelper(resp.StatusCode, "You have sent too many requests", &buf)}
	default:
		return errors.New(errStringHelper(resp.StatusCode, "API request returned an error", &buf))
	}
}

func errStringHelper(statusCode int, msg string, errBody *[]byte) string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(statusCode))
	buf.WriteString(": ")
	buf.WriteString(msg)
	buf.WriteString(", Body: ")
	buf.Write(*errBody)
	return buf.String()
}
