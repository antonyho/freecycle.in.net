package models

type Request struct {
	ItemId           uint
	ContactMethod    string
	Contact          string
	Email            string // mandatory for communication between owner and requestor
	ProposedDelivery string // proposed delivery method
	Offer            string // optional if any
}
