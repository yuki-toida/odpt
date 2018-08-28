package tran

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/tran"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/cache"
)

type UseCase struct {
	db  *gorm.DB
	cuc *cache.UseCase
}

func NewUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		db:  r.DB,
		cuc: cache.NewUseCase(r.Cache),
	}
}

func (u *UseCase) GetUpdateTime() tran.AdminTranTime {
	var row tran.AdminTranTime
	u.db.First(&row, 1)
	return row
}

func (u *UseCase) UpdateTranAt(date time.Time) {
	var row = u.GetUpdateTime()
	row.TranAt = &date
	u.db.Save(&row)
}

func (u *UseCase) UpdateMasterAt(date time.Time) {
	var row = u.GetUpdateTime()
	row.MasterAt = &date
	u.db.Save(&row)
}

func (u *UseCase) GetStationTimetable(sameAs string) (master.StationTimetableMaster, error) {
	var row master.StationTimetableMaster
	err := u.db.Preload("Objects").Where(&master.StationTimetableMaster{Base: master.Base{SameAs: sameAs}}).First(&row).Error
	if err == nil {
		calendar, _ := u.cuc.GetCalendar(row.CalendarSameAs)
		operator, _ := u.cuc.GetOperator(row.OperatorSameAs)
		railDirection, _ := u.cuc.GetRailDirection(row.RailDirectionSameAs)
		station, _ := u.cuc.GetStation(row.StationSameAs)
		row.Calendar = calendar
		row.Operator = operator
		row.RailDirection = railDirection
		row.Station = station
	}
	return row, err
}

func (u *UseCase) GetStationTimetableObject(ID int) (master.StationTimetableMasterObject, error) {
	var row master.StationTimetableMasterObject
	err := u.db.Preload("DestinationStations").Where(&master.StationTimetableMasterObject{ID: ID}).First(&row).Error
	if err == nil {
		trainType, _ := u.cuc.GetTrainType(row.TrainTypeSameAs)
		trainNames := u.cuc.GetStationTimetableObjectTrainNamesByObjectID(row.ID)
		viaRailways := u.cuc.GetStationTimetableObjectViaRailwaysByObjectID(row.ID)
		row.TrainType = trainType
		row.TrainNames = trainNames
		row.ViaRailways = viaRailways
	}
	return row, err
}

func (u *UseCase) GetTrainTimetable(trainSameAs string) (master.TrainTimetableMaster, error) {
	var row master.TrainTimetableMaster
	err := u.db.
		Preload("Objects").
		Preload("DestinationStations").
		Preload("DestinationStations.Station").
		Preload("OriginStations").
		Preload("OriginStations.Station").
		Where(&master.TrainTimetableMaster{TrainSameAs: trainSameAs}).
		First(&row).
		Error
	if err == nil {
		calendar, _ := u.cuc.GetCalendar(row.CalendarSameAs)
		operator, _ := u.cuc.GetOperator(row.OperatorSameAs)
		railDirection, _ := u.cuc.GetRailDirection(row.RailDirectionSameAs)
		trainType, _ := u.cuc.GetTrainType(row.TrainTypeSameAs)
		nexts := u.cuc.GetTrainTimetableNextsByID(row.ID)
		previous := u.cuc.GetTrainTimetablePreviousByID(row.ID)
		row.Calendar = calendar
		row.Operator = operator
		row.RailDirection = railDirection
		row.TrainType = trainType
		row.Nexts = nexts
		row.Previous = previous
	}
	return row, err
}

func (u *UseCase) GetTrains() []tran.TrainTran {
	var rows []tran.TrainTran
	u.db.Find(&rows)
	for i, v := range rows {
		railDirection, _ := u.cuc.GetRailDirection(v.RailDirectionSameAs)
		fromStation, _ := u.cuc.GetStation(v.FromStationSameAs)
		toStation, _ := u.cuc.GetStation(v.ToStationSameAs)
		rows[i].RailDirection = railDirection
		rows[i].FromStation = fromStation
		rows[i].ToStation = toStation
	}
	return rows
}

func (u *UseCase) GetTrain(sameAs string) (tran.TrainTran, error) {
	var row tran.TrainTran
	err := u.db.Preload("DestinationStations").Preload("OriginStations").Where(&tran.TrainTran{Base: tran.Base{SameAs: sameAs}}).First(&row).Error
	if err == nil {
		operator, _ := u.cuc.GetOperator(row.OperatorSameAs)
		railDirection, _ := u.cuc.GetRailDirection(row.RailDirectionSameAs)
		railway, _ := u.cuc.GetRailway(row.RailwaySameAs)
		trainType, _ := u.cuc.GetTrainType(row.TrainTypeSameAs)
		fromStation, _ := u.cuc.GetStation(row.FromStationSameAs)
		toStation, _ := u.cuc.GetStation(row.ToStationSameAs)
		row.Operator = operator
		row.RailDirection = railDirection
		row.Railway = railway
		row.TrainType = trainType
		row.FromStation = fromStation
		row.ToStation = toStation
	}
	return row, err
}

func (u *UseCase) GetTrainInformations() []tran.TrainInformationTran {
	var rows []tran.TrainInformationTran
	u.db.Preload("Railways").Find(&rows)
	for i, v := range rows {
		operator, _ := u.cuc.GetOperator(v.OperatorSameAs)
		railway, _ := u.cuc.GetRailway(v.RailwaySameAs)
		stationFrom, _ := u.cuc.GetStation(v.StationFromSameAs)
		stationTo, _ := u.cuc.GetStation(v.StationToSameAs)
		rows[i].Operator = operator
		rows[i].Railway = railway
		rows[i].StationFrom = stationFrom
		rows[i].StationTo = stationTo
	}
	return rows
}

func (u *UseCase) GetTrainInformation(sameAs string) (tran.TrainInformationTran, error) {
	var row tran.TrainInformationTran
	err := u.db.Preload("Railways").Where(&tran.TrainInformationTran{Base: tran.Base{SameAs: sameAs}}).First(&row).Error
	if err == nil {
		operator, _ := u.cuc.GetOperator(row.OperatorSameAs)
		railway, _ := u.cuc.GetRailway(row.RailwaySameAs)
		stationFrom, _ := u.cuc.GetStation(row.StationFromSameAs)
		stationTo, _ := u.cuc.GetStation(row.StationToSameAs)
		row.Operator = operator
		row.Railway = railway
		row.StationFrom = stationFrom
		row.StationTo = stationTo

		for i, v := range row.Railways {
			railway, _ := u.cuc.GetRailway(v.RailwaySameAs)
			row.Railways[i].Railway = railway
		}
	}
	return row, err
}
