package openweathermap

import (
	"fmt"
	"net/http"
	"testing"
)

var invalidOptionsTable = [][]Options{
	{WithLang("fk")},
	{WithLang("es"), WithUnit("celsuis")},
	{WithUnit("kelvin")},
	{WithHTTPClient(nil)},
	{nil, nil},
	{nil},
}

var validOptionsTable = [][]Options{
	{WithLang("es"), WithHTTPClient(http.DefaultClient)},
	{WithLang("sp"), WithUnit("F")},
	{},
}

func TestNewWithInvalidOptions(t *testing.T) {
	for _, opts := range invalidOptionsTable {
		_, err := New("apikey", opts...)
		if _, ok := err.(*IllegalArgumentError); ok {
			fmt.Printf("Received expected error message: %s\n", err.Error())
		} else {
			t.Error("Expected error, got none")
		}
	}
}

func TestNewWithValidOptions(t *testing.T) {
	for _, opts := range validOptionsTable {
		_, err := New("apikey", opts...)
		if err != nil {
			t.Errorf("Expected no errors. Got %v\n", err.Error())
		}
	}
}

func TestNewWithEmptyAPIKey(t *testing.T) {
	_, err := New("")
	if _, ok := err.(*IllegalArgumentError); ok {
		fmt.Printf("Received expected error message: %s\n", err.Error())
	} else {
		t.Error("Expected error, got none")
	}
}
