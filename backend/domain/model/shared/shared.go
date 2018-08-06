package shared

// Lang struct
type Lang struct {
	Ja string `json:"ja"`
	En string `json:"en"`
}

// Calendar struct
type Calendar struct {
	Context           string `json:"@context"`
	ID                string `json:"@id"`
	Type              string `json:"@type"`
	DcDate            string `json:"dc:date"`
	DcTitle           string `json:"dc:title"`
	OdptCalendarTitle Lang   `json:"odpt:calendarTitle"`
	OdptDayTitle      string `json:"odpt:day"`
	OdptDurationTitle string `json:"odpt:duration"`
	OwlSameAs         string `json:"owl:sameAs"`
}

// Operator struct
type Operator struct {
	Context           string `json:"@context"`
	ID                string `json:"@id"`
	Type              string `json:"@type"`
	DcDate            string `json:"dc:date"`
	DcTitle           string `json:"dc:title"`
	OdptOperatorTitle Lang   `json:"odpt:operatorTitle"`
	OwlSameAs         string `json:"owl:sameAs"`
}
