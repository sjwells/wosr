package main

type author struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ImageURL      string `json:"imageurl"`
	Biography     string `json:"biography"`
	TwitterHandle string `json:"twitterhandle"`
	TmeIdentifier string `json:"tmeidentifier"`
}
