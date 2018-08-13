package master

import "time"

// Base struct
type Base struct {
	ID      string
	SameAs  string     `gorm:"not null;unique"`
	Context string     `json:"-"`
	Type    string     `json:"-"`
	Date    *time.Time `gorm:"type:datetime" json:"-"`
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
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	Railways       []PassengerSurveyMasterRailway
	Stations       []PassengerSurveyMasterStation
	Objects        []PassengerSurveyMasterObject
}

// PassengerSurveyMasterRailway struct
type PassengerSurveyMasterRailway struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	RailwaySameAs           string
	Railway                 RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
}

// PassengerSurveyMasterStation struct
type PassengerSurveyMasterStation struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	StationSameAs           string
	Station                 StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationSameAs"`
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
	Color          string `json:"-"`
	Kana           string `json:"-"`
	LineCode       string
	OperatorSameAs string         `json:"-"`
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	RailwayTitleJa string
	RailwayTitleEn string
	Region         string `json:"-"`
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
	Title              string
	Lat                float64
	Long               float64
	OperatorSameAs     string         `json:"-"`
	Operator           OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	RailwaySameAs      string         `json:"-"`
	Railway            RailwayMaster  `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
	StationCode        string
	StationTitleJa     string
	StationTitleEn     string
	Region             string `json:"-"`
	ConnectingRailways []StationMasterConnectingRailway
	Exits              []StationMasterExit `json:"-"`
	PassengerSurveys   []StationMasterPassengerSurvey
	Timetables         []StationMasterTimetable
}

// StationMasterConnectingRailway struct
type StationMasterConnectingRailway struct {
	ID              uint   `json:"-"`
	StationMasterID string `json:"-"`
	RailwaySameAs   string
	Railway         *RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
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
	StationTimetable       StationTimetableMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationTimetableSameAs"`
}

// StationTimetableMaster struct
type StationTimetableMaster struct {
	Base
	CalendarSameAs      string              `json:"-"`
	Calendar            CalendarMaster      `gorm:"foreignkey:SameAs;association_foreignkey:CalendarSameAs"`
	NoteJa              string              `json:"-"`
	NoteEn              string              `json:"-"`
	OperatorSameAs      string              `json:"-"`
	RailDirectionSameAs string              `json:"-"`
	RailDirection       RailDirectionMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailDirectionSameAs"`
	RailwaySameAs       string              `json:"-"`
	RailwayTitleJa      string              `json:"-"`
	RailwayTitleEn      string              `json:"-"`
	Station             string
	StationTitleJa      string `json:"-"`
	StationTitleEn      string `json:"-"`
	Objects             []StationTimetableMasterObject
}

// StationTimetableMasterObject struct
type StationTimetableMasterObject struct {
	ID                       int
	StationTimetableMasterID string `json:"-"`
	ArrivalTime              string `json:"-"`
	CarComposition           int    `json:"-"`
	DepartureTime            string
	IsLast                   bool   `json:"-"`
	IsOrigin                 bool   `json:"-"`
	NoteJa                   string `json:"-"`
	NoteEn                   string `json:"-"`
	PlatformNameJa           string `json:"-"`
	PlatformNameEn           string `json:"-"`
	PlatformNumber           string `json:"-"`
	TrainSameAs              string
	TrainTypeSameAs          string          `json:"-"`
	TrainType                TrainTypeMaster `gorm:"foreignkey:SameAs;association_foreignkey:TrainTypeSameAs"`
	DestinationStations      []StationTimetableMasterObjectDestinationStation
	OriginStations           []StationTimetableMasterObjectOriginStation `json:"-"`
	TrainNames               []StationTimetableMasterObjectTrainName
	ViaRailways              []StationTimetableMasterObjectViaRailway
	ViaStations              []StationTimetableMasterObjectViaStation `json:"-"`
}

// StationTimetableMasterObjectDestinationStation struct
type StationTimetableMasterObjectDestinationStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

// StationTimetableMasterObjectOriginStation struct
type StationTimetableMasterObjectOriginStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

// StationTimetableMasterObjectTrainName struct
type StationTimetableMasterObjectTrainName struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	TrainNameJa                    string
	TrainNameEn                    string
}

// StationTimetableMasterObjectViaRailway struct
type StationTimetableMasterObjectViaRailway struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	RailwaySameAs                  string
}

// StationTimetableMasterObjectViaStation struct
type StationTimetableMasterObjectViaStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

// TrainTimetableMaster struct
type TrainTimetableMaster struct {
	Base
	CalendarSameAs      string
	NeedExtraFee        bool                `json:"-"`
	NoteJa              string              `json:"-"`
	NoteEn              string              `json:"-"`
	OperatorSameAs      string              `json:"-"`
	RailDirectionSameAs string              `json:"-"`
	RailDirection       RailDirectionMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailDirectionSameAs"`
	RailwaySameAs       string              `json:"-"`
	TrainSameAs         string
	TrainNumber         string
	TrainOwner          string
	TrainTypeSameAs     string          `json:"-"`
	TrainType           TrainTypeMaster `gorm:"foreignkey:SameAs;association_foreignkey:TrainTypeSameAs"`
}

// TrainTimetableMasterObject struct
type TrainTimetableMasterObject struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	ArrivalStation         string
	ArrivalTime            string
	DepartureStation       string
	DepartureTime          string
	NoteJa                 string `json:"-"`
	NoteEn                 string `json:"-"`
	PlatformNameJa         string `json:"-"`
	PlatformNameEn         string `json:"-"`
	PlatformNumber         string `json:"-"`
}

// TrainTimetableMasterDestinationStation struct
type TrainTimetableMasterDestinationStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

// TrainTimetableMasterOriginStation struct
type TrainTimetableMasterOriginStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

// TrainTimetableMasterNext struct
type TrainTimetableMasterNext struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainTimetableSameAs   string
}

// TrainTimetableMasterPrevious struct
type TrainTimetableMasterPrevious struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainTimetableSameAs   string
}

// TrainTimetableMasterTrainName struct
type TrainTimetableMasterTrainName struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainNameJa            string
	TrainNameEn            string
}

// TrainTimetableMasterViaRailway struct
type TrainTimetableMasterViaRailway struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	RailwaySameAs          string
}

// TrainTimetableMasterViaStation struct
type TrainTimetableMasterViaStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

// TrainTypeMaster struct
type TrainTypeMaster struct {
	Base
	Title            string
	OperatorSameAs   string
	Operator         OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	TrainTypeTitleJa string
	TrainTypeTitleEn string
}
