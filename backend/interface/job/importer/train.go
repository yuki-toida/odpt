package importer

import (
	"fmt"
	"sync"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/raw"
	"github.com/yuki-toida/refodpt/backend/domain/model/tran"
)

func (i Importer) passengerSurvey() {
	passengerSurveys, err := i.uc.GetPassengerSurveys()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PassengerSurvey]: %v\n", len(passengerSurveys))
	i.truncate("passenger_survey_masters")
	i.truncate("passenger_survey_master_railways")
	i.truncate("passenger_survey_master_stations")
	i.truncate("passenger_survey_master_objects")

	s := make([]interface{}, len(passengerSurveys))
	for i, v := range passengerSurveys {
		s[i] = v
	}
	async(s, i.createPassengerSurvey)
}

func (i Importer) createPassengerSurvey(passengerSurvey interface{}) {
	v, _ := passengerSurvey.(raw.PassengerSurvey)
	i.r.DB.Create(&master.PassengerSurveyMaster{
		Base: master.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
		},
		OperatorSameAs: v.OdptOperator,
	})

	for _, object := range v.OdptPassengerSurveyObject {
		i.r.DB.Create(&master.PassengerSurveyMasterObject{
			PassengerSurveyMasterID: v.ID,
			PassengerJourneys:       object.OdptPassengerJourneys,
			SurveyYear:              object.OdptSurveyYear,
		})
	}
	for _, railway := range v.OdptRailway {
		i.r.DB.Create(&master.PassengerSurveyMasterRailway{
			PassengerSurveyMasterID: v.ID,
			RailwaySameAs:           railway,
		})
	}
	for _, station := range v.OdptStation {
		i.r.DB.Create(&master.PassengerSurveyMasterStation{
			PassengerSurveyMasterID: v.ID,
			StationSameAs:           station,
		})
	}
}

func (i Importer) railDirection() {
	railDirections, err := i.uc.GetRailDirections()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[RailDirection]: %v\n", len(railDirections))
	i.truncate("rail_direction_masters")

	for _, v := range railDirections {
		i.r.DB.Create(&master.RailDirectionMaster{
			Base: master.Base{
				ID:      v.ID,
				SameAs:  v.OwlSameAs,
				Context: v.Context,
				Type:    v.Type,
				Date:    parseDate(v.DcDate),
			},
			Title:                v.DcTitle,
			RailDirectionTitleJa: v.OdptRailDirectionTitle.Ja,
			RailDirectionTitleEn: v.OdptRailDirectionTitle.En,
		})
	}
}

func (i Importer) railway() {
	railways, err := i.uc.GetRailways()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Railway]: %v\n", len(railways))
	i.truncate("railway_masters")
	i.truncate("railway_master_station_orders")

	s := make([]interface{}, len(railways))
	for i, v := range railways {
		s[i] = v
	}
	async(s, i.createRailway)
}

func (i Importer) createRailway(railway interface{}) {
	v, _ := railway.(raw.Railway)
	i.r.DB.Create(&master.RailwayMaster{
		Base: master.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
		},
		Title:          v.DcTitle,
		Color:          v.OdptColor,
		Kana:           v.OdptKana,
		LineCode:       v.OdptLineCode,
		OperatorSameAs: v.OdptOperator,
		RailwayTitleJa: v.OdptRailwayTitle.Ja,
		RailwayTitleEn: v.OdptRailwayTitle.En,
		Region:         v.UgRegion,
	})

	for _, order := range v.OdptStationOrder {
		i.r.DB.Create(&master.RailwayMasterStationOrder{
			RailwayMasterID: v.ID,
			Index:           order.OdptIndex,
			StationSameAs:   order.OdptStation,
			StationTitleJa:  order.OdptStationTitle.Ja,
			StationTitleEn:  order.OdptStationTitle.En,
		})
	}
}

func (i Importer) railwayFare() {
	railwayFares, err := i.uc.GetRailwayFares()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[RailwayFare]: %v\n", len(railwayFares))
	i.truncate("railway_fare_masters")

	s := make([]interface{}, len(railwayFares))
	for i, v := range railwayFares {
		s[i] = v
	}
	async(s, i.createRailwayFare)
}

func (i Importer) createRailwayFare(railwayFare interface{}) {
	v, _ := railwayFare.(raw.RailwayFare)
	i.r.DB.Create(&master.RailwayFareMaster{
		Base: master.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
		},
		OperatorSameAs:    v.OdptOperator,
		FromStationSameAs: v.OdptFromStation,
		ToStationSameAs:   v.OdptToStation,
		IcCardFare:        v.OdptIcCardFare,
		TicketFare:        v.OdptTicketFare,
		ChildIcCardFare:   v.OdptChildIcCardFare,
		ChildTicketFare:   v.OdptChildTicketFare,
		TicketType:        v.OdptTicketType,
	})
}

func (i Importer) station() {
	stations, err := i.uc.GetStations()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Station]: %v\n", len(stations))
	i.truncate("station_masters")
	i.truncate("station_master_connecting_railways")
	i.truncate("station_master_exits")
	i.truncate("station_master_passenger_surveys")
	i.truncate("station_master_timetables")

	s := make([]interface{}, len(stations))
	for i, v := range stations {
		s[i] = v
	}
	async(s, i.createStation)
}

func (i Importer) createStation(station interface{}) {
	v, _ := station.(raw.Station)
	i.r.DB.Create(&master.StationMaster{
		Base: master.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
		},
		Title:          v.DcTitle,
		Lat:            v.GeoLat,
		Long:           v.GeoLong,
		OperatorSameAs: v.OdptOperator,
		RailwaySameAs:  v.OdptRailway,
		StationCode:    v.OdptStationCode,
		StationTitleJa: v.OdptStationTitle.Ja,
		StationTitleEn: v.OdptStationTitle.En,
		Region:         v.UgRegion,
	})

	for _, x := range v.OdptConnectingRailway {
		i.r.DB.Create(&master.StationMasterConnectingRailway{
			StationMasterID: v.ID,
			RailwaySameAs:   x,
		})
	}
	for _, x := range v.OdptExit {
		i.r.DB.Create(&master.StationMasterExit{
			StationMasterID: v.ID,
			Exit:            x,
		})
	}
	for _, x := range v.OdptPassengerSurvey {
		i.r.DB.Create(&master.StationMasterPassengerSurvey{
			StationMasterID:       v.ID,
			PassengerSurveySameAs: x,
		})
	}
	for _, x := range v.OdptStationTimetable {
		i.r.DB.Create(&master.StationMasterTimetable{
			StationMasterID:        v.ID,
			StationTimetableSameAs: x,
		})
	}
}

func (i Importer) stationTimetable() {
	stationTimetables, err := i.uc.GetStationTimetables()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[StationTimetable]: %v\n", len(stationTimetables))
	i.truncate("station_timetable_masters")
	i.truncate("station_timetable_master_objects")
	i.truncate("station_timetable_master_object_destination_stations")
	i.truncate("station_timetable_master_object_origin_stations")
	i.truncate("station_timetable_master_object_train_names")
	i.truncate("station_timetable_master_object_via_railways")
	i.truncate("station_timetable_master_object_via_stations")

	for _, v := range stationTimetables {
		if v.ID != "" {
			i.r.DB.Create(&master.StationTimetableMaster{
				Base: master.Base{ID: v.ID,
					SameAs:  v.OwlSameAs,
					Context: v.Context,
					Type:    v.Type,
					Date:    parseDate(v.DcDate),
				},
				CalendarSameAs:      v.OdptCalendar,
				NoteJa:              v.OdptNote.Ja,
				NoteEn:              v.OdptNote.En,
				OperatorSameAs:      v.OdptOperator,
				RailDirectionSameAs: v.OdptRailDirection,
				RailwaySameAs:       v.OdptRailway,
				RailwayTitleJa:      v.OdptRailwayTitle.Ja,
				RailwayTitleEn:      v.OdptRailwayTitle.En,
				StationSameAs:       v.OdptStation,
				StationTitleJa:      v.OdptStationTitle.Ja,
				StationTitleEn:      v.OdptStationTitle.En,
			})

			ch := make(chan int, config.Config.DB.Pool/2)
			wg := sync.WaitGroup{}
			for index, object := range v.OdptStationTimetableObject {
				ch <- 1
				wg.Add(1)
				go func(id string, index int, object raw.StationTimetableObject) {
					i.createStationObject(id, index, object)
					<-ch
					wg.Done()
				}(v.ID, index, object)
			}
			wg.Wait()
		}
	}
}

func (i Importer) createStationObject(id string, index int, object raw.StationTimetableObject) {
	index++
	i.r.DB.Create(&master.StationTimetableMasterObject{
		StationTimetableMasterID: id,
		ArrivalTime:              object.OdptArrivalTime,
		CarComposition:           object.OdptCarComposition,
		DepartureTime:            object.OdptDepartureTime,
		IsLast:                   object.OdptIsLast,
		IsOrigin:                 object.OdptIsOrigin,
		NoteJa:                   object.OdptNote.Ja,
		NoteEn:                   object.OdptNote.En,
		PlatformNameJa:           object.OdptPlatformName.Ja,
		PlatformNameEn:           object.OdptPlatformName.En,
		PlatformNumber:           object.OdptPlatformNumber,
		TrainSameAs:              object.OdptTrain,
		TrainNumber:              object.OdptTrainNumber,
		TrainOwner:               object.OdptTrainOwner,
		TrainTypeSameAs:          object.OdptTrainType,
	})

	for _, x := range object.OdptDestinationStation {
		i.r.DB.Create(&master.StationTimetableMasterObjectDestinationStation{
			StationTimetableMasterID:       id,
			StationTimetableMasterObjectID: index,
			StationSameAs:                  x,
		})
	}
	for _, x := range object.OdptOriginStation {
		i.r.DB.Create(&master.StationTimetableMasterObjectOriginStation{
			StationTimetableMasterID:       id,
			StationTimetableMasterObjectID: index,
			StationSameAs:                  x,
		})
	}
	for _, x := range object.OdptTrainName {
		i.r.DB.Create(&master.StationTimetableMasterObjectTrainName{
			StationTimetableMasterID:       id,
			StationTimetableMasterObjectID: index,
			TrainNameJa:                    x.Ja,
			TrainNameEn:                    x.En,
		})
	}
	for _, x := range object.OdptViaRailway {
		i.r.DB.Create(&master.StationTimetableMasterObjectViaRailway{
			StationTimetableMasterID:       id,
			StationTimetableMasterObjectID: index,
			RailwaySameAs:                  x,
		})
	}
	for _, x := range object.OdptViaStation {
		i.r.DB.Create(&master.StationTimetableMasterObjectViaStation{
			StationTimetableMasterID:       id,
			StationTimetableMasterObjectID: index,
			StationSameAs:                  x,
		})
	}

}

func (i Importer) train() {
	trains, err := i.uc.GetTrains()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Train]: %v\n", len(trains))
	i.truncate("train_trans")
	i.truncate("train_tran_destination_stations")
	i.truncate("train_tran_origin_stations")
	i.truncate("train_tran_train_names")
	i.truncate("train_tran_via_railways")
	i.truncate("train_tran_via_stations")

	s := make([]interface{}, len(trains))
	for i, v := range trains {
		s[i] = v
	}
	async(s, i.createTrain)
}

func (i Importer) createTrain(train interface{}) {
	v, _ := train.(raw.Train)
	i.r.DB.Create(&tran.TrainTran{
		Base: tran.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
			Valid:   parseDate(v.DcValid),
		},
		CarComposition:      v.OdptCarComposition,
		Delay:               v.OdptDelay,
		FromStationSameAs:   v.OdptFromStation,
		Index:               v.OdptIndex,
		NoteJa:              v.OdptNote.Ja,
		NoteEn:              v.OdptNote.En,
		OperatorSameAs:      v.OdptOperator,
		RailDirectionSameAs: v.OdptRailDirection,
		RailwaySameAs:       v.OdptRailway,
		ToStationSameAs:     v.OdptToStation,
		TrainNumber:         v.OdptTrainNumber,
		TrainOwner:          v.OdptTrainOwner,
		TrainTypeSameAs:     v.OdptTrainType,
	})

	for _, x := range v.OdptDestinationStation {
		i.r.DB.Create(&tran.TrainTranDestinationStation{
			TrainTranID:   v.ID,
			StationSameAs: x,
		})
	}
	for _, x := range v.OdptOriginStation {
		i.r.DB.Create(&tran.TrainTranOriginStation{
			TrainTranID:   v.ID,
			StationSameAs: x,
		})
	}
	for _, x := range v.OdptTrainName {
		i.r.DB.Create(&tran.TrainTranTrainName{
			TrainTranID: v.ID,
			TrainNameJa: x.Ja,
			TrainNameEn: x.En,
		})
	}
	for _, x := range v.OdptViaRailway {
		i.r.DB.Create(&tran.TrainTranViaRailway{
			TrainTranID:   v.ID,
			RailwaySameAs: x,
		})
	}
	for _, x := range v.OdptViaStation {
		i.r.DB.Create(&tran.TrainTranViaStation{
			TrainTranID:   v.ID,
			StationSameAs: x,
		})
	}
}

func (i Importer) trainInformation() {
	trainInformations, err := i.uc.GetTrainInformations()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[TrainInformation]: %v\n", len(trainInformations))
	i.truncate("train_information_trans")
	i.truncate("train_information_tran_railways")

	s := make([]interface{}, len(trainInformations))
	for i, v := range trainInformations {
		s[i] = v
	}
	async(s, i.createTrainInformation)
}

func (i Importer) createTrainInformation(trainInformation interface{}) {
	v, _ := trainInformation.(raw.TrainInformation)
	i.r.DB.Create(&tran.TrainInformationTran{
		Base: tran.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
			Valid:   parseDate(v.DcValid),
		},
		OperatorSameAs:           v.OdptOperator,
		RailwaySameAs:            v.OdptRailway,
		ResumeEstimate:           v.OdptResumeEstimate,
		StationFromSameAs:        v.OdptStationFrom,
		StationToSameAs:          v.OdptStationTo,
		TimeOfOrigin:             parseDate(v.OdptTimeOfOrigin),
		TrainInformationAreaJa:   v.OdptTrainInformationArea.Ja,
		TrainInformationAreaEn:   v.OdptTrainInformationArea.En,
		TrainInformationCauseJa:  v.OdptTrainInformationCause.Ja,
		TrainInformationCauseEn:  v.OdptTrainInformationCause.En,
		TrainInformationKindJa:   v.OdptTrainInformationKind.Ja,
		TrainInformationKindEn:   v.OdptTrainInformationKind.En,
		TrainInformationLineJa:   v.OdptTrainInformationLine.Ja,
		TrainInformationLineEn:   v.OdptTrainInformationLine.En,
		TrainInformationRangeJa:  v.OdptTrainInformationRange.Ja,
		TrainInformationRangeEn:  v.OdptTrainInformationRange.En,
		TrainInformationStatusJa: v.OdptTrainInformationStatus.Ja,
		TrainInformationStatusEn: v.OdptTrainInformationStatus.En,
		TrainInformationTextJa:   v.OdptTrainInformationText.Ja,
		TrainInformationTextEn:   v.OdptTrainInformationText.En,
	})

	for _, x := range v.OdptRailways {
		i.r.DB.Create(&tran.TrainInformationTranRailway{
			TrainInformationTranID: v.ID,
			RailwaySameAs:          x,
		})
	}
}

func (i Importer) trainTimetable() {
	trainTimetables, err := i.uc.GetTrainTimetables()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[TrainTimetable]: %v\n", len(trainTimetables))
	i.truncate("train_timetable_masters")
	i.truncate("train_timetable_master_objects")
	i.truncate("train_timetable_master_destination_stations")
	i.truncate("train_timetable_master_origin_stations")
	i.truncate("train_timetable_master_nexts")
	i.truncate("train_timetable_master_previous")
	i.truncate("train_timetable_master_train_names")
	i.truncate("train_timetable_master_via_railways")
	i.truncate("train_timetable_master_via_stations")

	s := make([]interface{}, len(trainTimetables))
	for i, v := range trainTimetables {
		s[i] = v
	}
	async(s, i.createTrainTimetable)
}

func (i Importer) createTrainTimetable(trainTimetable interface{}) {
	v, _ := trainTimetable.(raw.TrainTimetable)
	i.r.DB.Create(&master.TrainTimetableMaster{
		Base: master.Base{
			ID:      v.ID,
			SameAs:  v.OwlSameAs,
			Context: v.Context,
			Type:    v.Type,
			Date:    parseDate(v.DcDate),
		},
		CalendarSameAs:      v.OdptCalendar,
		NeedExtraFee:        v.OdptNeedExtraFee,
		NoteJa:              v.OdptNote.Ja,
		NoteEn:              v.OdptNote.En,
		OperatorSameAs:      v.OdptOperator,
		RailDirectionSameAs: v.OdptRailDirection,
		RailwaySameAs:       v.OdptRailway,
		TrainSameAs:         v.OdptTrain,
		TrainNumber:         v.OdptTrainNumber,
		TrainOwner:          v.OdptTrainNumber,
		TrainTypeSameAs:     v.OdptTrainType,
	})

	for _, x := range v.OdptTrainTimetableObject {
		i.r.DB.Create(&master.TrainTimetableMasterObject{
			TrainTimetableMasterID: v.ID,
			ArrivalStation:         x.OdptArrivalStation,
			ArrivalTime:            x.OdptArrivalTime,
			DepartureStation:       x.OdptDepartureStation,
			DepartureTime:          x.OdptDepartureTime,
			NoteJa:                 x.OdptNote.Ja,
			NoteEn:                 x.OdptNote.En,
			PlatformNameJa:         x.OdptPlatformName.Ja,
			PlatformNameEn:         x.OdptPlatformName.En,
			PlatformNumber:         x.OdptPlatformNumber,
		})
	}
	for _, x := range v.OdptDestinationStation {
		i.r.DB.Create(&master.TrainTimetableMasterDestinationStation{
			TrainTimetableMasterID: v.ID,
			StationSameAs:          x,
		})
	}
	for _, x := range v.OdptOriginStation {
		i.r.DB.Create(&master.TrainTimetableMasterOriginStation{
			TrainTimetableMasterID: v.ID,
			StationSameAs:          x,
		})
	}
	for _, x := range v.OdptNextTrainTimetable {
		i.r.DB.Create(&master.TrainTimetableMasterNext{
			TrainTimetableMasterID: v.ID,
			TrainTimetableSameAs:   x,
		})
	}
	for _, x := range v.OdptPreviousTrainTimetable {
		i.r.DB.Create(&master.TrainTimetableMasterPrevious{
			TrainTimetableMasterID: v.ID,
			TrainTimetableSameAs:   x,
		})
	}
	for _, x := range v.OdptTrainName {
		i.r.DB.Create(&master.TrainTimetableMasterTrainName{
			TrainTimetableMasterID: v.ID,
			TrainNameJa:            x.Ja,
			TrainNameEn:            x.En,
		})
	}
	for _, x := range v.OdptViaRailway {
		i.r.DB.Create(&master.TrainTimetableMasterViaRailway{
			TrainTimetableMasterID: v.ID,
			RailwaySameAs:          x,
		})
	}
	for _, x := range v.OdptViaStation {
		i.r.DB.Create(&master.TrainTimetableMasterViaStation{
			TrainTimetableMasterID: v.ID,
			StationSameAs:          x,
		})
	}
}

func (i Importer) trainType() {
	trainTypes, err := i.uc.GetTrainTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[TrainType]: %v\n", len(trainTypes))
	i.truncate("train_type_masters")

	for _, v := range trainTypes {
		i.r.DB.Create(&master.TrainTypeMaster{
			Base: master.Base{
				ID:      v.ID,
				SameAs:  v.OwlSameAs,
				Context: v.Context,
				Type:    v.Type,
				Date:    parseDate(v.DcDate),
			},
			Title:            v.DcTitle,
			OperatorSameAs:   v.OdptOperator,
			TrainTypeTitleJa: v.OdptTrainTypeTitle.Ja,
			TrainTypeTitleEn: v.OdptTrainTypeTitle.En,
		})
	}
}
