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
		Preload("Railways").
		Preload("Stations").
		Preload("Objects").
		Where(&master.PassengerSurveyMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error
	return row, err
}

// GetRailwayMasters func
func (u *UseCase) GetRailwayMasters() []master.RailwayMaster {
	var rows []master.RailwayMaster
	u.Repository.DB.Preload("Operator").Find(&rows)
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

// GetStationMasters func
func (u *UseCase) GetStationMasters() []master.StationMaster {
	var rows []master.StationMaster
	u.Repository.DB.
		Preload("Operator").
		Preload("ConnectingRailways").
		Preload("Exits").
		Preload("PassengerSurveys").
		Preload("Timetables").
		Find(&rows)

	return rows
}

// GetStationMaster func
func (u *UseCase) GetStationMaster(sameAs string) (master.StationMaster, error) {
	var row master.StationMaster
	err := u.Repository.DB.
		Preload("Operator").
		Preload("ConnectingRailways").
		Preload("ConnectingRailways.Railway").
		Preload("Exits").
		Preload("PassengerSurveys").
		Preload("Timetables").
		Where(&master.StationMaster{Base: master.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, err
}
