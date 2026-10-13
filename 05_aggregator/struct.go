package main;

import "time"

type IncomingRawData struct {
	Timestamp 	time.Time 	`json:"timestamp"`
	Level 		string 		`json:"level"`
	Service 	string 		`json:"service"`
	Message		string 		`json:"message"` 
}

