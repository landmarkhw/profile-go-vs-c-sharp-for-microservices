package models

// Address represents a physical location.
type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
}

// Phone represents a phone number.
type Phone struct {
	Number    string
	PhoneType string
}

// Person contains details about an individual person.
type Person struct {
	ID              int32
	FirstName       string
	LastName        string
	Supervisor      *Person
	CareerCounselor *Person
	Phones          []Phone
	Emails          []string
	Addresses       []Address
}
