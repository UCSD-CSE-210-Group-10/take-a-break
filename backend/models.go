package main

type event struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Venue  string  `json:"venue"`
    Date   string  `json:"date"`
    Time   string  `json:"time"`
    Description string `json:"description"`
    Tags   string `json:"tags"`
    ImagePath  string `json:"imagepath"`
    Host   string `json:"host"`
    Contact string `json:"contact"`
}