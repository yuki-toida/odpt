package master

import "time"

// Base struct
type Base struct {
	ID      string     `json:"-"`
	SameAs  string     `gorm:"not null;unique"`
	Context string     `json:"-"`
	Type    string     `json:"-"`
	Date    *time.Time `gorm:"type:datetime" json:"-"`
}

// CategoryMaster struct
type CategoryMaster struct {
	ID        uint
	Category  string
	Type      string
	Name      string
	Desc      string
	UpdatedAt time.Time `gorm:"type:datetime" json:"-"`
}

// CalendarMaster struct
type CalendarMaster struct {
	Base
	Title           string
	CalendarTitleJa string
	CalendarTitleEn string
	Duration        string
	Days            []CalendarMasterDay
}

// CalendarMasterDay struct
type CalendarMasterDay struct {
	ID               uint   `json:"-"`
	CalendarMasterID string `json:"-"`
	Day              string
}

// OperatorMaster struct
type OperatorMaster struct {
	Base
	Title           string
	OperatorTitleJa string
	OperatorTitleEn string
}

// PassengerSurveyMaster struct
type PassengerSurveyMaster struct {
	Base
	OperatorSameAs string         `json:"-"`
	Operator       OperatorMaster `gorm:"auto_preload;foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	Railways       []PassengerSurveyMasterRailway
	Stations       []PassengerSurveyMasterStation
	Objects        []PassengerSurveyMasterObject
}

// PassengerSurveyMasterRailway struct
type PassengerSurveyMasterRailway struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	RailwaySameAs           string
}

// PassengerSurveyMasterStation struct
type PassengerSurveyMasterStation struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	StationSameAs           string
}

// PassengerSurveyMasterObject struct
type PassengerSurveyMasterObject struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	PassengerJourneys       int
	SurveyYear              int
}

// RailDirectionMaster struct
type RailDirectionMaster struct {
	Base
	Title                string
	RailDirectionTitleJa string
	RailDirectionTitleEn string
}

// RailwayMaster struct
type RailwayMaster struct {
	Base
	Title          string
	Color          string
	Kana           string
	LineCode       string
	OperatorSameAs string         `json:"-"`
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	RailwayTitleJa string
	RailwayTitleEn string
	Region         string
	StationOrders  []RailwayMasterStationOrder
}

// RailwayMasterStationOrder struct
type RailwayMasterStationOrder struct {
	ID              uint   `json:"-"`
	RailwayMasterID string `json:"-"`
	Index           int
	StationSameAs   string
	StationTitleJa  string
	StationTitleEn  string
}

// RailwayFareMaster struct
type RailwayFareMaster struct {
	Base
	OperatorSameAs    string         `json:"-"`
	Operator          OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	FromStationSameAs string
	ToStationSameAs   string
	IcCardFare        int
	TicketFare        int
	ChildIcCardFare   int
	ChildTicketFare   int
	TicketType        string
}

// StationMaster struct
type StationMaster struct {
	Base
	Lat                float64
	Long               float64
	OperatorSameAs     string         `json:"-"`
	Operator           OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	Railway            string
	StationCode        string
	StationTitleJa     string
	StationTitleEn     string
	Region             string
	ConnectingRailways []StationMasterConnectingRailway
	Exits              []StationMasterExit
	PassengerSurveys   []StationMasterPassengerSurvey
	Timetables         []StationMasterTimetable
}

// StationMasterConnectingRailway struct
type StationMasterConnectingRailway struct {
	ID              uint   `json:"-"`
	StationMasterID string `json:"-"`
	RailwaySameAs   string
	Railway         RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
}

// StationMasterExit struct
type StationMasterExit struct {
	ID              uint   `json:"-"`
	StationMasterID string `json:"-"`
	Exit            string
}

// StationMasterPassengerSurvey struct
type StationMasterPassengerSurvey struct {
	ID                    uint   `json:"-"`
	StationMasterID       string `json:"-"`
	PassengerSurveySameAs string
}

// StationMasterTimetable struct
type StationMasterTimetable struct {
	ID                     uint   `json:"-"`
	StationMasterID        string `json:"-"`
	StationTimetableSameAs string
}

// StationTimetable struct
type StationTimetable struct {
	ID             string
	SameAs         string `gorm:"not null;unique"`
	Context        string
	Type           string
	Date           *time.Time `gorm:"type:datetime"`
	Calendar       string
	NoteJa         string
	NoteEn         string
	OperatorSameAs string
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	RailDirection  string
	RailwaySameAs  string
	RailwayTitleJa string
	RailwayTitleEn string
	Station        string
	StationTitleJa string
	StationTitleEn string
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
}

// StationTimetableObjectDestinationStation struct
type StationTimetableObjectDestinationStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
}

// StationTimetableObjectOriginStation struct
type StationTimetableObjectOriginStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
}

// StationTimetableObjectViaRailway struct
type StationTimetableObjectViaRailway struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	RailwaySameAs            string
}

// StationTimetableObjectViaStation struct
type StationTimetableObjectViaStation struct {
	StationTimetableID       string `gorm:"primary_key"`
	StationTimetableObjectID int    `gorm:"primary_key;type:int"`
	ID                       int    `gorm:"primary_key;type:int"`
	StationSameAs            string
}

// TrainTimetable struct
type TrainTimetable struct {
	ID             string
	SameAs         string `gorm:"not null;unique"`
	Context        string
	Type           string
	Date           *time.Time `gorm:"type:datetime"`
	Calendar       string
	NeedExtraFee   bool
	NoteJa         string
	NoteEn         string
	OperatorSameAs string
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	RailDirection  string
	Railway        string
	Train          string
	TrainNameJa    string
	TrainNameEn    string
	TrainNumber    string
	TrainType      string
}

// TrainTimetableObject struct
type TrainTimetableObject struct {
	TrainTimetableID string `gorm:"primary_key"`
	ID               int    `gorm:"primary_key;type:int"`
	ArrivalStation   string
	ArrivalTime      string
	DepartureStation string
	DepartureTime    string
	NoteJa           string
	NoteEn           string
	PlatformNameJa   string
	PlatformNameEn   string
	PlatformNumber   string
}

// TrainTimetableDestinationStation struct
type TrainTimetableDestinationStation struct {
	TrainTimetableID string `gorm:"primary_key"`
	ID               int    `gorm:"primary_key;type:int"`
	StationSameAs    string
}

// TrainTimetableOriginStation struct
type TrainTimetableOriginStation struct {
	TrainTimetableID string `gorm:"primary_key"`
	ID               int    `gorm:"primary_key;type:int"`
	StationSameAs    string
}

// TrainTimetableNext struct
type TrainTimetableNext struct {
	TrainTimetableID     string `gorm:"primary_key"`
	ID                   int    `gorm:"primary_key;type:int"`
	TrainTimetableSameAs string
}

// TrainTimetablePrevious struct
type TrainTimetablePrevious struct {
	TrainTimetableID     string `gorm:"primary_key"`
	ID                   int    `gorm:"primary_key;type:int"`
	TrainTimetableSameAs string
}

// TrainTimetableViaRailway struct
type TrainTimetableViaRailway struct {
	TrainTimetableID string `gorm:"primary_key"`
	ID               int    `gorm:"primary_key;type:int"`
	RailwaySameAs    string
}

// TrainTimetableViaStation struct
type TrainTimetableViaStation struct {
	TrainTimetableID string `gorm:"primary_key"`
	ID               int    `gorm:"primary_key;type:int"`
	StationSameAs    string
}

// TrainType struct
type TrainType struct {
	ID               string
	SameAs           string `gorm:"not null;unique"`
	Context          string
	Type             string
	Date             *time.Time `gorm:"type:datetime"`
	Title            string
	OperatorSameAs   string
	Operator         OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	TrainTypeTitleJa string
	TrainTypeTitleEn string
}
