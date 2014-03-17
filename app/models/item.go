package models

type Item struct {
	Title         string
	Description   string
	Tags          string // seperated by comma, should be mapped with Tag model in database
	Condition     uint8  // percentage of status after depreciation
	Status        string // ([A]vailable / [H]old / [N]ot available...will be hidden) only avabile when it is posted by registered user
	Location      string // current location, free text, can be just district
	Handover      string // ([F]ace / [D]elivery) handover method
	Delivery      string // ([P]ost / [C]ourrier) delivery method if not handover
	ContactMethod string // free text for user
	Contact       string
	Email         string // optional for user notification purpose
	Owner         uint   // owner ID if it is registered (optional subscription)
}
