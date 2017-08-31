package main

import (
	"encoding/json"
)

type Options struct {
	Name string `json:"name"`
}

func NewOptions(s string) (Options, error) {
	o := Options{}
	err := json.Unmarshal([]byte(s), &o)
	return o, err
}
