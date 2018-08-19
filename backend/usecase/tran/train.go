package tran

import (
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

func (u *UseCase) GetStationTimetableMaster(sameAs string) (master.StationTimetableMaster, error) {
	var row master.StationTimetableMaster
	for _, v := range u.cuc.GetStationTimetables() {
		if v.SameAs == sameAs {
			row = v
			break
		}
	}

	var objects []master.StationTimetableMasterObject
	err := u.db.Where(&master.StationTimetableMasterObject{StationTimetableMasterID: row.ID}).Find(&objects).Error
	row.Objects = objects
	return row, err
}

func (u *UseCase) GetStationTimetableMasterObject(ID int) (master.StationTimetableMasterObject, error) {
	var row master.StationTimetableMasterObject
	err := u.db.Preload("TrainType").
		Preload("DestinationStations").
		Preload("TrainNames").
		Preload("ViaRailways").
		Where(&master.StationTimetableMasterObject{ID: ID}).
		First(&row).
		Error

	return row, err
}

func (u *UseCase) GetTrains() []tran.TrainTran {
	var rows []tran.TrainTran
	u.db.Preload("FromStation").Preload("RailDirection").Preload("ToStation").Find(&rows)
	return rows
}

func (u *UseCase) GetTrain(sameAs string) (tran.TrainTran, error) {
	var row tran.TrainTran
	err := u.db.
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

func (u *UseCase) GetTrainTimetable(trainSameAs string) (master.TrainTimetableMaster, error) {
	var row master.TrainTimetableMaster
	err := u.db.
		Preload("RailDirection").
		Preload("Objects").
		Preload("DestinationStations").
		Preload("OriginStations").
		Preload("Nexts").
		Preload("Previous").
		Preload("TrainNames").
		Preload("ViaRailways").
		Preload("ViaStations").
		Where(&master.TrainTimetableMaster{TrainSameAs: trainSameAs}).
		First(&row).
		Error

	return row, err
}

func (u *UseCase) GetTrainInformations() []tran.TrainInformationTran {
	var rows []tran.TrainInformationTran
	u.db.Preload("Railway").Preload("StationFrom").Preload("StationTo").Find(&rows)
	return rows
}

func (u *UseCase) GetTrainInformation(sameAs string) (tran.TrainInformationTran, error) {
	var row tran.TrainInformationTran
	err := u.db.
		Preload("Railway").
		Preload("StationFrom").
		Preload("StationTo").
		Preload("Railways").
		Where(&tran.TrainInformationTran{Base: tran.Base{SameAs: sameAs}}).
		First(&row).
		Error

	return row, err
}
