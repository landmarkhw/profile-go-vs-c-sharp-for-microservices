package models

// Address represents a physical location.
type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

// Phone represents a phone number.
type Phone struct {
	Number    string `json:"number"`
	PhoneType string `json:"phoneType"`
}

// Person contains details about an individual person.
type Person struct {
	ID              int64     `json:"id"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Supervisor      *Person   `json:"supervisor"`
	CareerCounselor *Person   `json:"careerCounselor"`
	Phones          []Phone   `json:"phones"`
	Emails          []string  `json:"emails"`
	Addresses       []Address `json:"addresses"`
}
