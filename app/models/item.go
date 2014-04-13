package models

type Item struct {
	Title         string
	Description   string
	Tags          string // seperated by comma, should be mapped with Tag model in database
	Condition     uint8  // percentage of status after depreciation
	Status        string // ([P]ending / [A]vailable / [H]old / [N]ot available...will be hidden) only avabile when it is posted by registered user
	Duration      int    // list duration. default 7 days
	Location      string // pickup location, free text, can be just district
	Handover      string // ([F]ace / [D]elivery) handover method
	Delivery      string // ([P]ost / [C]ourrier) delivery method if not handover
	ContactMethod string // free text for user
	Contact       string
	Email         string // mandatory. used for sending verification email. updating item status by replying email. contacting with request.
	Owner         uint   // owner ID if it is registered (optional subscription)
	PostDate      int64  // returned by time.Time.Now() unixepoch
	UpdateDate    int64  // returned by time.Time.Now() unixepoch
}
