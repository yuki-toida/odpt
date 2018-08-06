package train

import "github.com/yuki-toida/refodpt/backend/domain/model/shared"

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
	Context                string      `json:"@context"`
	ID                     string      `json:"@id"`
	Type                   string      `json:"@type"`
	DcDate                 string      `json:"dc:date"`
	DcTitle                string      `json:"dc:title"`
	OdptOperator           string      `json:"odpt:operator"`
	OdptRailDirectionTitle shared.Lang `json:"odpt:railDirectionTitle"`
	OwlSameAs              string      `json:"owl:sameAs"`
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
	OdptRainwayTitle shared.Lang    `json:"odpt:railwayTitle"`
	OdptStationOrder []StationOrder `json:"odpt:stationOrder"`
	OwlSameAs        string         `json:"owl:sameAs"`
	UgRegion         string         `json:"ug:region"`
}

// StationOrder struct
type StationOrder struct {
	OdptIndex        int         `json:"odpt:index"`
	OdptStation      string      `json:"odpt:station"`
	OdptStationTitle shared.Lang `json:"odpt:stationTitle"`
}
