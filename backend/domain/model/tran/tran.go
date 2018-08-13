package tran

import (
	"time"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
)

// Base struct
type Base struct {
	ID      string
	SameAs  string     `gorm:"not null;unique"`
	Context string     `json:"-"`
	Type    string     `json:"-"`
	Date    *time.Time `gorm:"type:datetime" json:"-"`
	Valid   *time.Time `gorm:"type:datetime" json:"-"`
}

// TrainTran struct
type TrainTran struct {
	Base
	CarComposition      int
	Delay               int
	FromStationSameAs   string               `json:"-"`
	FromStation         master.StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:FromStationSameAs"`
	Index               int
	NoteJa              string `json:"-"`
	NoteEn              string `json:"-"`
	OperatorSameAs      string
	RailDirectionSameAs string                     `json:"-"`
	RailDirection       master.RailDirectionMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailDirectionSameAs"`
	RailwaySameAs       string
	ToStationSameAs     string               `json:"-"`
	ToStation           master.StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:ToStationSameAs"`
	TrainNumber         string
	TrainOwner          string
	TrainTypeSameAs     string
	DestinationStations []TrainTranDestinationStation
	OriginStations      []TrainTranOriginStation
	TrainNames          []TrainTranTrainName  `json:"-"`
	ViaRailways         []TrainTranViaRailway `json:"-"`
	ViaStations         []TrainTranViaStation `json:"-"`
}

// TrainTranDestinationStation struct
type TrainTranDestinationStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}

// TrainTranOriginStation struct
type TrainTranOriginStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}

// TrainTranTrainName struct
type TrainTranTrainName struct {
	ID          uint   `json:"-"`
	TrainTranID string `json:"-"`
	TrainNameJa string
	TrainNameEn string
}

// TrainTranViaRailway struct
type TrainTranViaRailway struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	RailwaySameAs string
}

// TrainTranViaStation struct
type TrainTranViaStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}
