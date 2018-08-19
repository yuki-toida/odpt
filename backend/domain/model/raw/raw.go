package raw

type Lang struct {
	Ja string `json:"ja"`
	En string `json:"en"`
}

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

type Operator struct {
	Context           string `json:"@context"`
	ID                string `json:"@id"`
	Type              string `json:"@type"`
	DcDate            string `json:"dc:date"`
	DcTitle           string `json:"dc:title"`
	OdptOperatorTitle Lang   `json:"odpt:operatorTitle"`
	OwlSameAs         string `json:"owl:sameAs"`
}

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

type PassengerSurveyObject struct {
	OdptPassengerJourneys int `json:"odpt:passengerJourneys"`
	OdptSurveyYear        int `json:"odpt:surveyYear"`
}

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
	OdptRailwayTitle Lang           `json:"odpt:railwayTitle"`
	OdptStationOrder []StationOrder `json:"odpt:stationOrder"`
	OwlSameAs        string         `json:"owl:sameAs"`
	UgRegion         string         `json:"ug:region"`
}

type StationOrder struct {
	OdptIndex        int    `json:"odpt:index"`
	OdptStation      string `json:"odpt:station"`
	OdptStationTitle Lang   `json:"odpt:stationTitle"`
}

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

type Station struct {
	Context               string   `json:"@context"`
	ID                    string   `json:"@id"`
	Type                  string   `json:"@type"`
	DcDate                string   `json:"dc:date"`
	DcTitle               string   `json:"dc:title"`
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
	OdptTrainName          []Lang   `json:"odpt:trainName"`
	OdptTrainNumber        string   `json:"odpt:trainNumber"`
	OdptTrainOwner         string   `json:"odpt:trainOwner"`
	OdptTrainType          string   `json:"odpt:trainType"`
	OdptViaRailway         []string `json:"odpt:viaRailway"`
	OdptViaStation         []string `json:"odpt:viaStation"`
}

type Train struct {
	Context                string   `json:"@context"`
	ID                     string   `json:"@id"`
	Type                   string   `json:"@type"`
	DcDate                 string   `json:"dc:date"`
	DcValid                string   `json:"dc:valid"`
	OdptCarComposition     int      `json:"odpt:carComposition"`
	OdptDelay              int      `json:"odpt:delay"`
	OdptDestinationStation []string `json:"odpt:destinationStation"`
	OdptFromStation        string   `json:"odpt:fromStation"`
	OdptIndex              int      `json:"odpt:index"`
	OdptNote               Lang     `json:"odpt:note"`
	OdptOperator           string   `json:"odpt:operator"`
	OdptOriginStation      []string `json:"odpt:originStation"`
	OdptRailDirection      string   `json:"odpt:railDirection"`
	OdptRailway            string   `json:"odpt:railway"`
	OdptToStation          string   `json:"odpt:toStation"`
	OdptTrainName          []Lang   `json:"odpt:trainName"`
	OdptTrainNumber        string   `json:"odpt:trainNumber"`
	OdptTrainOwner         string   `json:"odpt:trainOwner"`
	OdptTrainType          string   `json:"odpt:trainType"`
	OdptViaRailway         []string `json:"odpt:viaRailway"`
	OdptViaStation         []string `json:"odpt:viaStation"`
	OwlSameAs              string   `json:"owl:sameAs"`
}

type TrainInformation struct {
	Context                    string   `json:"@context"`
	ID                         string   `json:"@id"`
	Type                       string   `json:"@type"`
	DcDate                     string   `json:"dc:date"`
	DcValid                    string   `json:"dc:valid"`
	OdptOperator               string   `json:"odpt:operator"`
	OdptRailway                string   `json:"odpt:railway"`
	OdptResumeEstimate         string   `json:"odpt:resumeEstimate"`
	OdptStationFrom            string   `json:"odpt:stationFrom"`
	OdptStationTo              string   `json:"odpt:stationTo"`
	OdptTimeOfOrigin           string   `json:"odpt:timeOfOrigin"`
	OdptTrainInformationArea   Lang     `json:"odpt:trainInformationArea"`
	OdptTrainInformationCause  Lang     `json:"odpt:trainInformationCauseTitle"`
	OdptTrainInformationKind   Lang     `json:"odpt:trainInformationKind"`
	OdptTrainInformationLine   Lang     `json:"odpt:trainInformationLineTitle"`
	OdptTrainInformationRange  Lang     `json:"odpt:trainInformationRange"`
	OdptTrainInformationStatus Lang     `json:"odpt:trainInformationStatus"`
	OdptTrainInformationText   Lang     `json:"odpt:trainInformationText"`
	OdptRailways               []string `json:"odpt:transferRailways"`
	OwlSameAs                  string   `json:"owl:sameAs"`
}

type TrainTimetable struct {
	Context                    string                 `json:"@context"`
	ID                         string                 `json:"@id"`
	Type                       string                 `json:"@type"`
	DcDate                     string                 `json:"dc:date"`
	OdptCalendar               string                 `json:"odpt:calendar"`
	OdptDestinationStation     []string               `json:"odpt:destinationStation"`
	OdptNeedExtraFee           bool                   `json:"odpt:needExtraFee"`
	OdptNextTrainTimetable     []string               `json:"odpt:nextTrainTimetable"`
	OdptNote                   Lang                   `json:"odpt:note"`
	OdptOperator               string                 `json:"odpt:operator"`
	OdptOriginStation          []string               `json:"odpt:originStation"`
	OdptPreviousTrainTimetable []string               `json:"odpt:previousTrainTimetable"`
	OdptRailDirection          string                 `json:"odpt:railDirection"`
	OdptRailway                string                 `json:"odpt:railway"`
	OdptTrain                  string                 `json:"odpt:train"`
	OdptTrainName              []Lang                 `json:"odpt:trainName"`
	OdptTrainNumber            string                 `json:"odpt:trainNumber"`
	OdptTrainTimetableObject   []TrainTimetableObject `json:"odpt:trainTimetableObject"`
	OdptTrainType              string                 `json:"odpt:trainType"`
	OdptViaRailway             []string               `json:"odpt:viaRailway"`
	OdptViaStation             []string               `json:"odpt:viaStation"`
	OwlSameAs                  string                 `json:"owl:sameAs"`
}

type TrainTimetableObject struct {
	OdptArrivalStation   string `json:"odpt:arrivalStation"`
	OdptArrivalTime      string `json:"odpt:arrivalTime"`
	OdptDepartureStation string `json:"odpt:departureStation"`
	OdptDepartureTime    string `json:"odpt:departureTime"`
	OdptNote             Lang   `json:"odpt:note"`
	OdptPlatformName     Lang   `json:"odpt:platformName"`
	OdptPlatformNumber   string `json:"odpt:platformNumber"`
}

type TrainType struct {
	Context            string `json:"@context"`
	ID                 string `json:"@id"`
	Type               string `json:"@type"`
	DcDate             string `json:"dc:date"`
	DcTitle            string `json:"dc:title"`
	OdptOperator       string `json:"odpt:operator"`
	OdptTrainTypeTitle Lang   `json:"odpt:trainTypeTitle"`
	OwlSameAs          string `json:"owl:sameAs"`
}
