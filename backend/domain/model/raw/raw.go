package raw

// Lang struct
type Lang struct {
	Ja string `json:"ja"`
	En string `json:"en"`
}

// Calendar struct
type Calendar struct {
	Context           string   `json:"@context"`
	ID                string   `json:"@id"`
	Type              string   `json:"@type"`
	DcDate            string   `json:"dc:date"`
	DcTitle           string   `json:"dc:title"`
	OdptCalendarTitle Lang     `json:"odpt:calendarTitle"`
	OdptDay           []string `json:"odpt:day"`
	OdptDuration      string   `json:"odpt:duration"`
	OwlSameAs         string   `json:"owl:sameAs"`
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

// PassengerSurvey struct
type PassengerSurvey struct {
	Context                   string                  `json:"@context"`
	ID                        string                  `json:"@id"`
	Type                      string                  `json:"@type"`
	DcDate                    string                  `json:"dc:date"`
	OdptOperator              string                  `json:"odpt:operator"`
	OdptPassengerSurveyObject []PassengerSurveyObject `json:"odpt:passengerSurveyObject"`
	OdptRailway               []string                `json:"odpt:railway"`
	OdptStation               []string                `json:"odpt:station"`
	OwlSameAs                 string                  `json:"owl:sameAs"`
}

// PassengerSurveyObject struct
type PassengerSurveyObject struct {
	OdptPassengerJourneys int `json:"odpt:passengerJourneys"`
	OdptSurveyYear        int `json:"odpt:surveyYear"`
}

// RailDirection struct
type RailDirection struct {
	Context                string `json:"@context"`
	ID                     string `json:"@id"`
	Type                   string `json:"@type"`
	DcDate                 string `json:"dc:date"`
	DcTitle                string `json:"dc:title"`
	OdptOperator           string `json:"odpt:operator"`
	OdptRailDirectionTitle Lang   `json:"odpt:railDirectionTitle"`
	OwlSameAs              string `json:"owl:sameAs"`
}

// Railway struct
type Railway struct {
	Context          string         `json:"@context"`
	ID               string         `json:"@id"`
	Type             string         `json:"@type"`
	DcDate           string         `json:"dc:date"`
	DcTitle          string         `json:"dc:title"`
	OdptColor        string         `json:"odpt:color"`
	OdptKana         string         `json:"odpt:kana"`
	OdptLineCode     string         `json:"odpt:lineCode"`
	OdptOperator     string         `json:"odpt:operator"`
	OdptRainwayTitle Lang           `json:"odpt:railwayTitle"`
	OdptStationOrder []StationOrder `json:"odpt:stationOrder"`
	OwlSameAs        string         `json:"owl:sameAs"`
	UgRegion         string         `json:"ug:region"`
}

// StationOrder struct
type StationOrder struct {
	OdptIndex        int    `json:"odpt:index"`
	OdptStation      string `json:"odpt:station"`
	OdptStationTitle Lang   `json:"odpt:stationTitle"`
}

// RailwayFare struct
type RailwayFare struct {
	Context             string `json:"@context"`
	ID                  string `json:"@id"`
	Type                string `json:"@type"`
	DcDate              string `json:"dc:date"`
	OdptChildIcCardFare int    `json:"odpt:childIcCardFare"`
	OdptChildTicketFare int    `json:"odpt:childTicketFare"`
	OdptFromStation     string `json:"odpt:fromStation"`
	OdptIcCardFare      int    `json:"odpt:icCardFare"`
	OdptOperator        string `json:"odpt:operator"`
	OdptTicketFare      int    `json:"odpt:ticketFare"`
	OdptTicketType      string `json:"odpt:ticketType"`
	OdptToStation       string `json:"odpt:toStation"`
	OwlSameAs           string `json:"owl:sameAs"`
}

// Station struct
type Station struct {
	Context               string   `json:"@context"`
	ID                    string   `json:"@id"`
	Type                  string   `json:"@type"`
	DcDate                string   `json:"dc:date"`
	GeoLat                float64  `json:"geo:lat"`
	GeoLong               float64  `json:"geo:long"`
	OdptConnectingRailway []string `json:"odpt:connectingRailway"`
	OdptExit              []string `json:"odpt:exit"`
	OdptOperator          string   `json:"odpt:operator"`
	OdptPassengerSurvey   []string `json:"odpt:passengerSurvey"`
	OdptRailway           string   `json:"odpt:railway"`
	OdptStationCode       string   `json:"odpt:stationCode"`
	OdptStationTimetable  []string `json:"odpt:stationTimetable"`
	OdptStationTitle      Lang     `json:"odpt:stationTitle"`
	OwlSameAs             string   `json:"owl:sameAs"`
	UgRegion              string   `json:"ug:region"`
}

// StationTimetable struct
type StationTimetable struct {
	Context                    string                   `json:"@context"`
	ID                         string                   `json:"@id"`
	Type                       string                   `json:"@type"`
	DcDate                     string                   `json:"dc:date"`
	OdptCalendar               string                   `json:"odpt:calendar"`
	OdptNote                   Lang                     `json:"odpt:note"`
	OdptOperator               string                   `json:"odpt:operator"`
	OdptRailDirection          string                   `json:"odpt:railDirection"`
	OdptRailway                string                   `json:"odpt:railway"`
	OdptRailwayTitle           Lang                     `json:"odpt:railwayTitle"`
	OdptStation                string                   `json:"odpt:station"`
	OdptStationTimetableObject []StationTimetableObject `json:"odpt:stationTimetableObject"`
	OdptStationTitle           Lang                     `json:"odpt:stationTitle"`
	OwlSameAs                  string                   `json:"owl:sameAs"`
}

// StationTimetableObject struct
type StationTimetableObject struct {
	OdptArrivalTime        string   `json:"odpt:arrivalTime"`
	OdptCarComposition     int      `json:"odpt:carComposition"`
	OdptDepartureTime      string   `json:"odpt:departureTime"`
	OdptDestinationStation []string `json:"odpt:destinationStation"`
	OdptIsLast             bool     `json:"odpt:isLast"`
	OdptIsOrigin           bool     `json:"odpt:isOrigin"`
	OdptNote               Lang     `json:"odpt:note"`
	OdptOriginStation      []string `json:"odpt:originStation"`
	OdptPlatformName       Lang     `json:"odpt:platformName"`
	OdptPlatformNumber     string   `json:"odpt:platformNumber"`
	OdptTrain              string   `json:"odpt:train"`
	OdptTrainName          Lang     `json:"odpt:trainName"`
	OdptTrainType          string   `json:"odpt:trainType"`
	OdptViaRailway         []string `json:"odpt:viaRailway"`
	OdptViaStation         []string `json:"odpt:viaStation"`
}
