package train

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/tran"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/shared"
)

// UseCase type
type UseCase struct {
	Repository    *repository.Repository
	SharedUseCase *shared.UseCase
}

// NewUseCase func
func NewUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		Repository:    r,
		SharedUseCase: shared.NewUseCase(r),
	}
}

// GetPassengerSurveyMaster func
func (u *UseCase) GetPassengerSurveyMaster(sameAs string) (master.PassengerSurveyMaster, error) {
	var row master.PassengerSurveyMaster
	err := u.Repository.DB.
		Preload("Operator").
		Preload("Railways").
		Preload("Railways.Railway").
		Preload("Stations").
		Preload("Stations.Station").
		Preload("Objects").
		Where(&master.PassengerSurveyMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error
	return row, err
}

// GetRailwayMasters func
func (u *UseCase) GetRailwayMasters() []master.RailwayMaster {
	var rows []master.RailwayMaster
	u.Repository.DB.Preload("Operator").Preload("StationOrders").Find(&rows)
	return rows
}

// GetRailwayMaster func
func (u *UseCase) GetRailwayMaster(sameAs string) (master.RailwayMaster, error) {
	var row master.RailwayMaster
	error := u.Repository.DB.
		Preload("Operator").
		Preload("StationOrders").
		Where(&master.RailwayMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, error
}

// GetStationMaster func
func (u *UseCase) GetStationMaster(sameAs string) (master.StationMaster, error) {
	var row master.StationMaster
	err := u.Repository.DB.
		Preload("Operator").
		Preload("Railway").
		Preload("ConnectingRailways").
		Preload("ConnectingRailways.Railway").
		Preload("PassengerSurveys").
		Preload("Timetables").
		Preload("Timetables.StationTimetable").
		Where(&master.StationMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, err
}

// GetStationTimetableMaster func
func (u *UseCase) GetStationTimetableMaster(sameAs string) (master.StationTimetableMaster, error) {
	var row master.StationTimetableMaster
	err := u.Repository.DB.
		Preload("Calendar").
		Preload("RailDirection").
		Preload("Objects").
		Where(&master.StationTimetableMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, err
}

// GetStationTimetableObjectMaster func
func (u *UseCase) GetStationTimetableObjectMaster(ID int) (master.StationTimetableMasterObject, error) {
	var row master.StationTimetableMasterObject
	err := u.Repository.DB.
		Preload("TrainType").
		Preload("DestinationStations").
		Preload("TrainNames").
		Preload("ViaRailways").
		Where(&master.StationTimetableMasterObject{ID: ID}).
		First(&row).
		Error

	return row, err
}

// GetTrains func
func (u *UseCase) GetTrains() []tran.TrainTran {
	var rows []tran.TrainTran
	u.Repository.DB.
		Preload("FromStation").
		Preload("RailDirection").
		Preload("ToStation").
		Find(&rows)

	return rows
}

// GetTrain func
func (u *UseCase) GetTrain(sameAs string) (tran.TrainTran, error) {
	var row tran.TrainTran
	err := u.Repository.DB.
		Preload("FromStation").
		Preload("RailDirection").
		Preload("ToStation").
		Preload("DestinationStations").
		Preload("OriginStations").
		Where(&tran.TrainTran{Base: tran.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, err
}
