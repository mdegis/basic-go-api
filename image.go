package main

import "time"

type Image struct {
	Id        int       `json:"id"`
	Location  string    `json:"location"`
	Path      string    `json:"path"`
	Date      time.Time `json:"date"`
}

type Images []Image
