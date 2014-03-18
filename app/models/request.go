package models

type Request struct {
	ItemId           uint
	ContactMethod    string
	Contact          string
	Email            string // optional
	ProposedDelivery string // proposed delivery method
	Offer            string // optional if any
}
