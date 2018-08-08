package model

import "time"

// PassengerSurvey struct
type PassengerSurvey struct {
	ID        string `gorm:"primary_key"`
	SameAs    string
	Context   string
	Type      string
	Date      *time.Time `gorm:"type:datetime"`
	Operator  string
	UpdatedAt time.Time `gorm:"type:datetime" json:"-"`
}

// PassengerSurveyRailway struct
type PassengerSurveyRailway struct {
	PassengerSurveyID string    `gorm:"primary_key"`
	RailwaySameAs     string    `gorm:"primary_key"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"-"`
}

// PassengerSurveyStation struct
type PassengerSurveyStation struct {
	PassengerSurveyID string    `gorm:"primary_key"`
	StationSameAs     string    `gorm:"primary_key"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"-"`
}

// PassengerSurveyObject struct
type PassengerSurveyObject struct {
	PassengerSurveyID string `gorm:"primary_key"`
	ID                int    `gorm:"primary_key;type:int"`
	PassengerJourneys int
	SurveyYear        int
	UpdatedAt         time.Time `gorm:"type:datetime" json:"-"`
}

// RailDirection struct
type RailDirection struct {
	ID                   string `gorm:"primary_key"`
	SameAs               string
	Context              string
	Type                 string
	Date                 *time.Time `gorm:"type:datetime"`
	Title                string
	RailDirectionTitleJa string
	RailDirectionTitleEn string
	UpdatedAt            time.Time `gorm:"type:datetime" json:"-"`
}

// Railway struct
type Railway struct {
	ID             string `gorm:"primary_key"`
	SameAs         string
	Context        string
	Type           string
	Date           *time.Time `gorm:"type:datetime"`
	Title          string
	Color          string
	Kana           string
	LineCode       string
	Operator       string
	RainwayTitleJa string
	RainwayTitleEn string
	Region         string
	UpdatedAt      time.Time `gorm:"type:datetime" json:"-"`
}

// RailwayStationOrder struct
type RailwayStationOrder struct {
	RailwayID      string `gorm:"primary_key"`
	Index          int    `gorm:"primary_key;type:int"`
	StationSameAs  string
	StationTitleJa string
	StationTitleEn string
	UpdatedAt      time.Time `gorm:"type:datetime" json:"-"`
}

// RailwayFare struct
type RailwayFare struct {
	ID                string `gorm:"primary_key"`
	SameAs            string
	Context           string
	Type              string
	Date              *time.Time `gorm:"type:datetime"`
	Operator          string
	FromStationSameAs string
	ToStationSameAs   string
	IcCardFare        int
	TicketFare        int
	ChildIcCardFare   int
	ChildTicketFare   int
	TicketType        string
	UpdatedAt         time.Time `gorm:"type:datetime" json:"-"`
}

// Station struct
type Station struct {
	ID             string `gorm:"primary_key"`
	SameAs         string
	Context        string
	Type           string
	Date           *time.Time `gorm:"type:datetime"`
	Lat            float64
	Long           float64
	Operator       string
	Railway        string
	StationCode    string
	StationTitleJa string
	StationTitleEn string
	Region         string
	UpdatedAt      time.Time `gorm:"type:datetime" json:"-"`
}

// StationConnectingRailway struct
type StationConnectingRailway struct {
	StationID     string    `gorm:"primary_key"`
	RailwaySameAs string    `gorm:"primary_key"`
	UpdatedAt     time.Time `gorm:"type:datetime" json:"-"`
}

// StationExit struct
type StationExit struct {
	StationID string    `gorm:"primary_key"`
	Exit      string    `gorm:"primary_key"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"-"`
}

// StationPassengerSurvey struct
type StationPassengerSurvey struct {
	StationID             string    `gorm:"primary_key"`
	PassengerSurveySameAs string    `gorm:"primary_key"`
	UpdatedAt             time.Time `gorm:"type:datetime" json:"-"`
}

// StationStationTimetable struct
type StationStationTimetable struct {
	StationID              string    `gorm:"primary_key"`
	StationTimetableSameAs string    `gorm:"primary_key"`
	UpdatedAt              time.Time `gorm:"type:datetime" json:"-"`
}

// StationTimetable struct
type StationTimetable struct {
	ID             string `gorm:"primary_key"`
	SameAs         string
	Context        string
	Type           string
	Date           *time.Time
	Calendar       string
	NoteJa         string
	NoteEn         string
	Operator       string
	RailDirection  string
	RailwaySameAs  string
	RailwayTitleJa string
	RailwayTitleEn string
	Station        string
	StationTitleJa string
	StationTitleEn string
	UpdatedAt      time.Time
}

// StationTimetableObject struct
type StationTimetableObject struct {
	StationTimetableID string `gorm:"primary_key"`
	ID                 int    `gorm:"primary_key;type:int"`
	ArrivalTime        string
	CarComposition     int
	DepartureTime      string
	IsLast             bool
	IsOrigin           bool
	NoteJa             string
	NoteEn             string
	PlatformNameJa     string
	PlatformNameEn     string
	PlatformNumber     string
	Train              string
	TrainNameJa        string
	TrainNameEn        string
	TrainType          string
	UpdatedAt          time.Time `gorm:"type:datetime" json:"-"`
}

// StationTimetableObjectDestinationStation struct
type StationTimetableObjectDestinationStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
	UpdatedAt                time.Time `gorm:"type:datetime" json:"-"`
}

// StationTimetableObjectOriginStation struct
type StationTimetableObjectOriginStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
	UpdatedAt                time.Time `gorm:"type:datetime" json:"-"`
}

// StationTimetableObjectViaRailway struct
type StationTimetableObjectViaRailway struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	RailwaySameAs            string
	UpdatedAt                time.Time `gorm:"type:datetime" json:"-"`
}

// StationTimetableObjectViaStation struct
type StationTimetableObjectViaStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
	UpdatedAt                time.Time `gorm:"type:datetime" json:"-"`
}
