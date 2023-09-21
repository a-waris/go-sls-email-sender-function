package main

type TaggedOfferContactForm struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Tag        string   `json:"tag"`
	Tags       []string `json:"tags"`
	Additional string   `json:"additional"`
}
