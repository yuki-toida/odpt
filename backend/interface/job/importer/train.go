package importer

import (
	"fmt"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
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

	for _, v := range passengerSurveys {
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

	for _, v := range railways {
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
}

func (i Importer) railwayFare() {
	railwayFares, err := i.uc.GetRailwayFares()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[RailwayFare]: %v\n", len(railwayFares))
	i.truncate("railway_fare_masters")

	for _, v := range railwayFares {
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

	for _, v := range stations {
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
				Station:             v.OdptStation,
				StationTitleJa:      v.OdptStationTitle.Ja,
				StationTitleEn:      v.OdptStationTitle.En,
			})

			for index, object := range v.OdptStationTimetableObject {
				index++
				i.r.DB.Create(&master.StationTimetableMasterObject{
					StationTimetableMasterID: v.ID,
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
					Train:                    object.OdptTrain,
					TrainType:                object.OdptTrainType,
				})

				for _, x := range object.OdptDestinationStation {
					i.r.DB.Create(&master.StationTimetableMasterObjectDestinationStation{
						StationTimetableMasterID:       v.ID,
						StationTimetableMasterObjectID: index,
						StationSameAs:                  x,
					})
				}
				for _, x := range object.OdptOriginStation {
					i.r.DB.Create(&master.StationTimetableMasterObjectOriginStation{
						StationTimetableMasterID:       v.ID,
						StationTimetableMasterObjectID: index,
						StationSameAs:                  x,
					})
				}
				for _, x := range object.OdptTrainName {
					i.r.DB.Create(&master.StationTimetableMasterObjectTrainName{
						StationTimetableMasterID:       v.ID,
						StationTimetableMasterObjectID: index,
						TrainNameJa:                    x.Ja,
						TrainNameEn:                    x.En,
					})
				}
				for _, x := range object.OdptViaRailway {
					i.r.DB.Create(&master.StationTimetableMasterObjectViaRailway{
						StationTimetableMasterID:       v.ID,
						StationTimetableMasterObjectID: index,
						RailwaySameAs:                  x,
					})
				}
				for _, x := range object.OdptViaStation {
					i.r.DB.Create(&master.StationTimetableMasterObjectViaStation{
						StationTimetableMasterID:       v.ID,
						StationTimetableMasterObjectID: index,
						StationSameAs:                  x,
					})
				}
			}
		}
	}
}

func (i Importer) trainTimetable() {
	trainTimetables, err := i.uc.GetTrainTimetables()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[TrainTimetable]: %v\n", len(trainTimetables))
	i.r.DB.Delete(master.TrainTimetable{})
	i.r.DB.Delete(master.TrainTimetableObject{})
	i.r.DB.Delete(master.TrainTimetableDestinationStation{})
	i.r.DB.Delete(master.TrainTimetableOriginStation{})
	i.r.DB.Delete(master.TrainTimetableNext{})
	i.r.DB.Delete(master.TrainTimetablePrevious{})
	i.r.DB.Delete(master.TrainTimetableViaRailway{})
	i.r.DB.Delete(master.TrainTimetableViaStation{})

	for _, v := range trainTimetables {
		i.r.DB.Create(&master.TrainTimetable{
			ID:             v.ID,
			SameAs:         v.OwlSameAs,
			Context:        v.Context,
			Type:           v.Type,
			Date:           parseDate(v.DcDate),
			Calendar:       v.OdptCalendar,
			NeedExtraFee:   v.OdptNeedExtraFee,
			NoteJa:         v.OdptNote.Ja,
			NoteEn:         v.OdptNote.En,
			OperatorSameAs: v.OdptOperator,
			RailDirection:  v.OdptRailDirection,
			Railway:        v.OdptRailway,
			Train:          v.OdptTrain,
			TrainNameJa:    v.OdptTrainName.Ja,
			TrainNameEn:    v.OdptTrainName.En,
			TrainNumber:    v.OdptTrainNumber,
			TrainType:      v.OdptTrainType,
		})

		for index, object := range v.OdptTrainTimetableObject {
			i.r.DB.Create(&master.TrainTimetableObject{
				TrainTimetableID: v.ID,
				ID:               index + 1,
				ArrivalStation:   object.OdptArrivalStation,
				ArrivalTime:      object.OdptArrivalTime,
				DepartureStation: object.OdptDepartureStation,
				DepartureTime:    object.OdptDepartureTime,
				NoteJa:           object.OdptNote.Ja,
				NoteEn:           object.OdptNote.En,
				PlatformNameJa:   object.OdptPlatformName.Ja,
				PlatformNameEn:   object.OdptPlatformName.En,
				PlatformNumber:   object.OdptPlatformNumber,
			})
		}
		for index, x := range v.OdptDestinationStation {
			i.r.DB.Create(&master.TrainTimetableDestinationStation{
				TrainTimetableID: v.ID,
				ID:               index + 1,
				StationSameAs:    x,
			})
		}
		for index, x := range v.OdptOriginStation {
			i.r.DB.Create(&master.TrainTimetableOriginStation{
				TrainTimetableID: v.ID,
				ID:               index + 1,
				StationSameAs:    x,
			})
		}
		for index, x := range v.OdptNextTrainTimetable {
			i.r.DB.Create(&master.TrainTimetableNext{
				TrainTimetableID:     v.ID,
				ID:                   index + 1,
				TrainTimetableSameAs: x,
			})
		}
		for index, x := range v.OdptPreviousTrainTimetable {
			i.r.DB.Create(&master.TrainTimetablePrevious{
				TrainTimetableID:     v.ID,
				ID:                   index + 1,
				TrainTimetableSameAs: x,
			})
		}
		for index, x := range v.OdptViaRailway {
			i.r.DB.Create(&master.TrainTimetableViaRailway{
				TrainTimetableID: v.ID,
				ID:               index + 1,
				RailwaySameAs:    x,
			})
		}
		for index, x := range v.OdptViaStation {
			i.r.DB.Create(&master.TrainTimetableViaStation{
				TrainTimetableID: v.ID,
				ID:               index + 1,
				StationSameAs:    x,
			})
		}
	}
}

func (i Importer) trainType() {
	trainTypes, err := i.uc.GetTrainTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[TrainType]: %v\n", len(trainTypes))
	i.r.DB.Delete(master.TrainType{})

	for _, v := range trainTypes {
		i.r.DB.Create(&master.TrainType{
			ID:               v.ID,
			SameAs:           v.OwlSameAs,
			Context:          v.Context,
			Type:             v.Type,
			Date:             parseDate(v.DcDate),
			Title:            v.DcTitle,
			OperatorSameAs:   v.OdptOperator,
			TrainTypeTitleJa: v.OdptTrainTypeTitle.Ja,
			TrainTypeTitleEn: v.OdptTrainTypeTitle.En,
		})
	}
}
