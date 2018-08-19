package master

import "time"

type Base struct {
	ID      string
	SameAs  string     `gorm:"not null;unique"`
	Context string     `json:"-"`
	Type    string     `json:"-"`
	Date    *time.Time `gorm:"type:datetime" json:"-"`
}

type CalendarMaster struct {
	Base
	Title           string
	CalendarTitleJa string
	CalendarTitleEn string
	Duration        string
	Days            []CalendarMasterDay
}

type CalendarMasterDay struct {
	ID               uint   `json:"-"`
	CalendarMasterID string `json:"-"`
	Day              string
}

type OperatorMaster struct {
	Base
	Title           string
	OperatorTitleJa string
	OperatorTitleEn string
}

type PassengerSurveyMaster struct {
	Base
	OperatorSameAs string         `json:"-"`
	Operator       OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	Railways       []PassengerSurveyMasterRailway
	Stations       []PassengerSurveyMasterStation
	Objects        []PassengerSurveyMasterObject
}

type PassengerSurveyMasterRailway struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	RailwaySameAs           string
	Railway                 RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
}

type PassengerSurveyMasterStation struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	StationSameAs           string
	Station                 StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationSameAs"`
}

type PassengerSurveyMasterObject struct {
	ID                      uint   `json:"-"`
	PassengerSurveyMasterID string `json:"-"`
	PassengerJourneys       int
	SurveyYear              int
}

type RailDirectionMaster struct {
	Base
	Title                string
	RailDirectionTitleJa string
	RailDirectionTitleEn string
}

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

type RailwayMasterStationOrder struct {
	ID              uint   `json:"-"`
	RailwayMasterID string `json:"-"`
	Index           int
	StationSameAs   string
	StationTitleJa  string
	StationTitleEn  string
}

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

type StationMasterConnectingRailway struct {
	ID              uint   `json:"-"`
	StationMasterID string `json:"-"`
	RailwaySameAs   string
	Railway         *RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
}

type StationMasterExit struct {
	ID              uint   `json:"-"`
	StationMasterID string `json:"-"`
	Exit            string
}

type StationMasterPassengerSurvey struct {
	ID                    uint   `json:"-"`
	StationMasterID       string `json:"-"`
	PassengerSurveySameAs string
}

type StationMasterTimetable struct {
	ID                     uint   `json:"-"`
	StationMasterID        string `json:"-"`
	StationTimetableSameAs string
	StationTimetable       StationTimetableMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationTimetableSameAs"`
}

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
	StationSameAs       string
	StationTitleJa      string `json:"-"`
	StationTitleEn      string `json:"-"`
	Objects             []StationTimetableMasterObject
}

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
	TrainNumber              string
	TrainOwner               string
	TrainTypeSameAs          string                                           `json:"-"`
	TrainType                TrainTypeMaster                                  `gorm:"foreignkey:SameAs;association_foreignkey:TrainTypeSameAs"`
	DestinationStations      []StationTimetableMasterObjectDestinationStation `gorm:"foreignkey:StationTimetableMasterObjectID"`
	OriginStations           []StationTimetableMasterObjectOriginStation      `json:"-"`
	TrainNames               []StationTimetableMasterObjectTrainName
	ViaRailways              []StationTimetableMasterObjectViaRailway
	ViaStations              []StationTimetableMasterObjectViaStation `json:"-"`
}

type StationTimetableMasterObjectDestinationStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

type StationTimetableMasterObjectOriginStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

type StationTimetableMasterObjectTrainName struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	TrainNameJa                    string
	TrainNameEn                    string
}

type StationTimetableMasterObjectViaRailway struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	RailwaySameAs                  string
}

type StationTimetableMasterObjectViaStation struct {
	ID                             uint   `json:"-"`
	StationTimetableMasterID       string `json:"-"`
	StationTimetableMasterObjectID int    `json:"-"`
	StationSameAs                  string
}

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
	TrainOwner          string          `json:"-"`
	TrainTypeSameAs     string          `json:"-"`
	TrainType           TrainTypeMaster `gorm:"foreignkey:SameAs;association_foreignkey:TrainTypeSameAs"`
	Objects             []TrainTimetableMasterObject
	DestinationStations []TrainTimetableMasterDestinationStation
	OriginStations      []TrainTimetableMasterOriginStation
	Nexts               []TrainTimetableMasterNext
	Previous            []TrainTimetableMasterPrevious
	TrainNames          []TrainTimetableMasterTrainName `json:"-"`
	ViaRailways         []TrainTimetableMasterViaRailway
	ViaStations         []TrainTimetableMasterViaStation `json:"-"`
}

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

type TrainTimetableMasterDestinationStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

type TrainTimetableMasterOriginStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

type TrainTimetableMasterNext struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainTimetableSameAs   string
}

type TrainTimetableMasterPrevious struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainTimetableSameAs   string
}

type TrainTimetableMasterTrainName struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	TrainNameJa            string
	TrainNameEn            string
}

type TrainTimetableMasterViaRailway struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	RailwaySameAs          string
}

type TrainTimetableMasterViaStation struct {
	ID                     uint   `json:"-"`
	TrainTimetableMasterID string `json:"-"`
	StationSameAs          string
}

type TrainTypeMaster struct {
	Base
	Title            string
	OperatorSameAs   string
	Operator         OperatorMaster `gorm:"foreignkey:SameAs;association_foreignkey:OperatorSameAs"`
	TrainTypeTitleJa string
	TrainTypeTitleEn string
}
