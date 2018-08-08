package importer

import (
	"fmt"

	"github.com/yuki-toida/refodpt/backend/domain/model"
)

func (i Importer) passengerSurvey() {
	passengerSurveys, err := i.uc.GetPassengerSurvey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PassengerSurvey]: %v\n", len(passengerSurveys))
	i.r.Delete(model.PassengerSurvey{})
	i.r.Delete(model.PassengerSurveyObject{})
	i.r.Delete(model.PassengerSurveyRailway{})
	i.r.Delete(model.PassengerSurveyStation{})

	for _, v := range passengerSurveys {
		i.r.Create(model.PassengerSurvey{
			ID:        v.ID,
			SameAs:    v.OwlSameAs,
			Context:   v.Context,
			Type:      v.Type,
			Date:      parseDate(v.DcDate),
			Operator:  v.OdptOperator,
			UpdatedAt: i.now,
		})

		for index, object := range v.OdptPassengerSurveyObject {
			i.r.Create(model.PassengerSurveyObject{
				PassengerSurveyID: v.ID,
				ID:                index + 1,
				PassengerJourneys: object.OdptPassengerJourneys,
				SurveyYear:        object.OdptSurveyYear,
				UpdatedAt:         i.now,
			})
		}
		for _, railway := range v.OdptRailway {
			i.r.Create(model.PassengerSurveyRailway{
				PassengerSurveyID: v.ID,
				RailwaySameAs:     railway,
				UpdatedAt:         i.now,
			})
		}
		for _, station := range v.OdptStation {
			i.r.Create(model.PassengerSurveyStation{
				PassengerSurveyID: v.ID,
				StationSameAs:     station,
				UpdatedAt:         i.now,
			})
		}
	}
}

func (i Importer) railDirection() {
	railDirections, err := i.uc.GetRailDirection()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[RailDirection]: %v\n", len(railDirections))
	i.r.Delete(model.RailDirection{})

	for _, v := range railDirections {
		i.r.Create(model.RailDirection{
			ID:                   v.ID,
			SameAs:               v.OwlSameAs,
			Context:              v.Context,
			Type:                 v.Type,
			Date:                 parseDate(v.DcDate),
			Title:                v.DcTitle,
			RailDirectionTitleJa: v.OdptRailDirectionTitle.Ja,
			RailDirectionTitleEn: v.OdptRailDirectionTitle.En,
			UpdatedAt:            i.now,
		})
	}
}

func (i Importer) railway() {
	railways, err := i.uc.GetRailway()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Railway]: %v\n", len(railways))
	i.r.Delete(model.Railway{})
	i.r.Delete(model.RailwayStationOrder{})

	for _, v := range railways {
		i.r.Create(model.Railway{
			ID:             v.ID,
			SameAs:         v.OwlSameAs,
			Context:        v.Context,
			Type:           v.Type,
			Date:           parseDate(v.DcDate),
			Title:          v.DcTitle,
			Color:          v.OdptColor,
			Kana:           v.OdptKana,
			LineCode:       v.OdptLineCode,
			Operator:       v.OdptOperator,
			RainwayTitleJa: v.OdptRainwayTitle.Ja,
			RainwayTitleEn: v.OdptRainwayTitle.En,
			Region:         v.UgRegion,
			UpdatedAt:      i.now,
		})

		for _, order := range v.OdptStationOrder {
			i.r.Create(model.RailwayStationOrder{
				RailwayID:      v.ID,
				Index:          order.OdptIndex,
				StationSameAs:  order.OdptStation,
				StationTitleJa: order.OdptStationTitle.Ja,
				StationTitleEn: order.OdptStationTitle.En,
				UpdatedAt:      i.now,
			})
		}
	}
}

func (i Importer) railwayFare() {
	railwayFares, err := i.uc.GetRailwayFare()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[RailwayFare]: %v\n", len(railwayFares))
	i.r.Delete(model.RailwayFare{})

	for _, v := range railwayFares {
		i.r.Create(model.RailwayFare{
			ID:                v.ID,
			SameAs:            v.OwlSameAs,
			Context:           v.Context,
			Type:              v.Type,
			Date:              parseDate(v.DcDate),
			Operator:          v.OdptOperator,
			FromStationSameAs: v.OdptFromStation,
			ToStationSameAs:   v.OdptToStation,
			IcCardFare:        v.OdptIcCardFare,
			TicketFare:        v.OdptTicketFare,
			ChildIcCardFare:   v.OdptChildIcCardFare,
			ChildTicketFare:   v.OdptChildTicketFare,
			TicketType:        v.OdptTicketType,
			UpdatedAt:         i.now,
		})
	}
}

func (i Importer) station() {
	stations, err := i.uc.GetStation()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Station]: %v\n", len(stations))
	i.r.Delete(model.Station{})
	i.r.Delete(model.StationConnectingRailway{})
	i.r.Delete(model.StationExit{})
	i.r.Delete(model.StationPassengerSurvey{})
	i.r.Delete(model.StationTimetable{})

	for _, v := range stations {
		i.r.Create(model.Station{
			ID:             v.ID,
			SameAs:         v.OwlSameAs,
			Context:        v.Context,
			Type:           v.Type,
			Date:           parseDate(v.DcDate),
			Lat:            v.GeoLat,
			Long:           v.GeoLong,
			Operator:       v.OdptOperator,
			Railway:        v.OdptRailway,
			StationCode:    v.OdptStationCode,
			StationTitleJa: v.OdptStationTitle.Ja,
			StationTitleEn: v.OdptStationTitle.En,
			Region:         v.UgRegion,
			UpdatedAt:      i.now,
		})

		for _, connectingRailway := range v.OdptConnectingRailway {
			i.r.Create(model.StationConnectingRailway{
				StationID:     v.ID,
				RailwaySameAs: connectingRailway,
				UpdatedAt:     i.now,
			})
		}
		for _, exit := range v.OdptExit {
			i.r.Create(model.StationExit{
				StationID: v.ID,
				Exit:      exit,
				UpdatedAt: i.now,
			})
		}
		for _, passengerSurvey := range v.OdptPassengerSurvey {
			i.r.Create(model.StationPassengerSurvey{
				StationID:             v.ID,
				PassengerSurveySameAs: passengerSurvey,
				UpdatedAt:             i.now,
			})
		}
		for _, stationTimetable := range v.OdptStationTimetable {
			i.r.Create(model.StationStationTimetable{
				StationID:              v.ID,
				StationTimetableSameAs: stationTimetable,
				UpdatedAt:              i.now,
			})
		}
	}
}

func (i Importer) stationTimetable() {
	stationTimetables, err := i.uc.GetStationTimetable()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[StationTimetable]: %v\n", len(stationTimetables))
	i.r.Delete(model.StationTimetable{})
	i.r.Delete(model.StationTimetableObject{})
	i.r.Delete(model.StationTimetableObjectDestinationStation{})
	i.r.Delete(model.StationTimetableObjectOriginStation{})
	i.r.Delete(model.StationTimetableObjectViaRailway{})
	i.r.Delete(model.StationTimetableObjectViaStation{})

	for _, v := range stationTimetables {
		i.r.Create(model.StationTimetable{
			ID:             v.ID,
			SameAs:         v.OwlSameAs,
			Context:        v.Context,
			Type:           v.Type,
			Date:           parseDate(v.DcDate),
			Calendar:       v.OdptCalendar,
			NoteJa:         v.OdptNote.Ja,
			NoteEn:         v.OdptNote.En,
			Operator:       v.OdptOperator,
			RailDirection:  v.OdptRailDirection,
			RailwaySameAs:  v.OdptRailway,
			RailwayTitleJa: v.OdptRailwayTitle.Ja,
			RailwayTitleEn: v.OdptRailwayTitle.En,
			Station:        v.OdptStation,
			StationTitleJa: v.OdptStationTitle.Ja,
			StationTitleEn: v.OdptStationTitle.En,
			UpdatedAt:      i.now,
		})
		fmt.Printf("[ID]: %v\n", v.ID)

		for ix, object := range v.OdptStationTimetableObject {
			index := ix + 1
			i.r.Create(model.StationTimetableObject{
				StationTimetableID: v.ID,
				ID:                 index,
				ArrivalTime:        object.OdptArrivalTime,
				CarComposition:     object.OdptCarComposition,
				DepartureTime:      object.OdptDepartureTime,
				IsLast:             object.OdptIsLast,
				IsOrigin:           object.OdptIsOrigin,
				NoteJa:             object.OdptNote.Ja,
				NoteEn:             object.OdptNote.En,
				PlatformNameJa:     object.OdptPlatformName.Ja,
				PlatformNameEn:     object.OdptPlatformName.En,
				PlatformNumber:     object.OdptPlatformNumber,
				Train:              object.OdptTrain,
				TrainNameJa:        object.OdptTrainName.Ja,
				TrainNameEn:        object.OdptTrainName.En,
				TrainType:          object.OdptTrainType,
				UpdatedAt:          i.now,
			})

			for ids, ds := range object.OdptDestinationStation {
				i.r.Create(model.StationTimetableObjectDestinationStation{
					StationTimetableID:       v.ID,
					StationTimetableObjectID: index,
					ID:            ids + 1,
					StationSameAs: ds,
					UpdatedAt:     i.now,
				})
			}
			for ios, os := range object.OdptOriginStation {
				i.r.Create(model.StationTimetableObjectOriginStation{
					StationTimetableID:       v.ID,
					StationTimetableObjectID: index,
					ID:            ios + 1,
					StationSameAs: os,
					UpdatedAt:     i.now,
				})
			}
			for ivr, vr := range object.OdptViaRailway {
				i.r.Create(model.StationTimetableObjectViaRailway{
					StationTimetableID:       v.ID,
					StationTimetableObjectID: index,
					ID:            ivr + 1,
					RailwaySameAs: vr,
					UpdatedAt:     i.now,
				})
			}
			for ivs, vs := range object.OdptViaStation {
				i.r.Create(model.StationTimetableObjectViaStation{
					StationTimetableID:       v.ID,
					StationTimetableObjectID: index,
					ID:            ivs + 1,
					StationSameAs: vs,
					UpdatedAt:     i.now,
				})
			}
		}
	}
}
