package train

import (
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
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
