package main

import "time"

type Temperature struct {
	Time  time.Time `json:"time,omitempty"`
	Value float32   `json:"value,omitempty"`
}

type Ph struct {
	Time  time.Time `json:"time,omitempty"`
	Value float32   `json:"value,omitempty"`
}

type Measurement struct {
	Time  time.Time
	Temperature *Temperature `json:"temperature,omitempty"`
	Ph          *Ph          `json:"ph,omitempty"`
}
