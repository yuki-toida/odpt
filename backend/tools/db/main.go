package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/tran"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
	"github.com/yuki-toida/refodpt/backend/infrastructure/db"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/interface/job/importer"
)

func init() {
	config.Init("../../", "../../.env")
}

func main() {
	db := db.Connect()
	db.LogMode(true)
	db.AutoMigrate(
		master.CalendarMaster{},
		master.CalendarMasterDay{},
		master.OperatorMaster{},
		master.PassengerSurveyMaster{},
		master.PassengerSurveyMasterObject{},
		master.PassengerSurveyMasterRailway{},
		master.PassengerSurveyMasterStation{},
		master.RailDirectionMaster{},
		master.RailwayMaster{},
		master.RailwayMasterStationOrder{},
		master.RailwayFareMaster{},
		master.StationMaster{},
		master.StationMasterConnectingRailway{},
		master.StationMasterExit{},
		master.StationMasterPassengerSurvey{},
		master.StationMasterTimetable{},
		master.StationTimetableMaster{},
		master.StationTimetableMasterObject{},
		master.StationTimetableMasterObjectDestinationStation{},
		master.StationTimetableMasterObjectOriginStation{},
		master.StationTimetableMasterObjectTrainName{},
		master.StationTimetableMasterObjectViaRailway{},
		master.StationTimetableMasterObjectViaStation{},
		master.TrainTimetableMaster{},
		master.TrainTimetableMasterObject{},
		master.TrainTimetableMasterDestinationStation{},
		master.TrainTimetableMasterOriginStation{},
		master.TrainTimetableMasterNext{},
		master.TrainTimetableMasterPrevious{},
		master.TrainTimetableMasterTrainName{},
		master.TrainTimetableMasterViaRailway{},
		master.TrainTimetableMasterViaStation{},
		master.TrainTypeMaster{},
		tran.AdminTranTime{},
		tran.TrainTran{},
		tran.TrainTranDestinationStation{},
		tran.TrainTranOriginStation{},
		tran.TrainTranTrainName{},
		tran.TrainTranViaRailway{},
		tran.TrainTranViaStation{},
		tran.TrainInformationTran{},
		tran.TrainInformationTranRailway{},
	)

	cc := cache.NewCache()
	r := repository.NewRepository(db, cc)
	importer.NewImporter(r).Master()
}
