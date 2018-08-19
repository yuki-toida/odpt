package tran

import (
	"time"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
)

type Base struct {
	ID      string
	SameAs  string     `gorm:"not null;unique"`
	Context string     `json:"-"`
	Type    string     `json:"-"`
	Date    *time.Time `gorm:"type:datetime" json:"-"`
	Valid   *time.Time `gorm:"type:datetime" json:"-"`
}

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

type TrainTranDestinationStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}

type TrainTranOriginStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}

type TrainTranTrainName struct {
	ID          uint   `json:"-"`
	TrainTranID string `json:"-"`
	TrainNameJa string
	TrainNameEn string
}

type TrainTranViaRailway struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	RailwaySameAs string
}

type TrainTranViaStation struct {
	ID            uint   `json:"-"`
	TrainTranID   string `json:"-"`
	StationSameAs string
}

type TrainInformationTran struct {
	Base
	OperatorSameAs           string
	RailwaySameAs            string
	Railway                  master.RailwayMaster `gorm:"foreignkey:SameAs;association_foreignkey:RailwaySameAs"`
	ResumeEstimate           string
	StationFromSameAs        string
	StationFrom              master.StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationFromSameAs"`
	StationToSameAs          string
	StationTo                master.StationMaster `gorm:"foreignkey:SameAs;association_foreignkey:StationToSameAs"`
	TimeOfOrigin             *time.Time           `gorm:"type:datetime"`
	TrainInformationAreaJa   string
	TrainInformationAreaEn   string
	TrainInformationCauseJa  string
	TrainInformationCauseEn  string
	TrainInformationKindJa   string
	TrainInformationKindEn   string
	TrainInformationLineJa   string
	TrainInformationLineEn   string
	TrainInformationRangeJa  string
	TrainInformationRangeEn  string
	TrainInformationStatusJa string
	TrainInformationStatusEn string
	TrainInformationTextJa   string
	TrainInformationTextEn   string
	Railways                 []TrainInformationTranRailway
}

type TrainInformationTranRailway struct {
	ID                     uint   `json:"-"`
	TrainInformationTranID string `json:"-"`
	RailwaySameAs          string
}
