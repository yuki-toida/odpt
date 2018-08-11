package importer

import (
	"fmt"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
)

func (i Importer) category() {
	categories := []master.CategoryMaster{
		{Category: "共通", Type: "odpt:Calendar", Name: "曜日・日付区分", Desc: "曜日・日付区分が記述される", UpdatedAt: i.now},
		{Category: "共通", Type: "odpt:Operator", Name: "事業者", Desc: "公共交通機関の事業者が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:PassengerSurvey", Name: "駅別乗降人員", Desc: "駅の乗降数集計結果が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:RailDirection", Name: "進行方向", Desc: "上り、下りなど、列車の進行方向が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:Railway", Name: "路線情報", Desc: "鉄道における路線、運行系統が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:RailwayFare", Name: "運賃情報", Desc: "運賃が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:Station", Name: "駅情報", Desc: "駅の名称や緯度系など、駅に関する情報が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:StationTimetable", Name: "駅時刻表", Desc: "駅を発着する列車の時刻表情報が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:Train", Name: "列車情報", Desc: "列車の位置、行き先など、列車のリアルタイムな情報が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:TrainTimetable", Name: "列車時刻表", Desc: "列車がどの駅にいつ到着するか、出発するか記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:TrainInformation", Name: "運行情報", Desc: "鉄道路線のリアルタイムな運行状況が記述される", UpdatedAt: i.now},
		{Category: "鉄道", Type: "odpt:TrainType", Name: "列車種別", Desc: "普通、快速など、列車の種別が記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:Bus", Name: "バス情報", Desc: "バス車両の位置、行先など、バス車両のリアルタイムな情報が記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:BusTimetable", Name: "バス時刻表", Desc: "バスがあるバス停、バス標柱にいつ到着するか、いつ出発するかが記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:BusroutePattern", Name: "バス運行路線情報", Desc: "バス運行における運行路線情報が記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:BusroutePatternFare", Name: "バス運賃情報", Desc: "バス運賃が記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:BusstopPole", Name: "バス時刻表", Desc: "バスがあるバス停、バス標柱にいつ到着するか、いつ出発するかが記述される", UpdatedAt: i.now},
		{Category: "バス", Type: "odpt:BusstopPoleTimetable", Name: "バス停標柱時刻表", Desc: "バスがあるバス停標柱にいつ到着するか、いつ出発するかが記述される", UpdatedAt: i.now},
		{Category: "航空機", Type: "odpt:FlightInfomationArrival", Name: "フライト出発情報", Desc: "空港に当日到着する航空機のリアルタイムな情報が記述される", UpdatedAt: i.now},
		{Category: "航空機", Type: "odpt:FlightInformationDeparture", Name: "フライト出発情報", Desc: "空港に当日到着する航空機のリアルタイムな情報が記述される", UpdatedAt: i.now},
		{Category: "航空機", Type: "odpt:FlightSchedule", Name: "航空機月間時刻表", Desc: "航空機の予定される発着時刻情報が記述される", UpdatedAt: i.now},
	}

	i.truncate("category_masters")
	for _, v := range categories {
		i.r.DB.Create(&v)
	}
}

func (i Importer) calendar() {
	calendars, err := i.uc.GetCalendars()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Calendars]: %v\n", len(calendars))
	i.truncate("calendar_masters")
	i.truncate("calendar_master_days")

	for _, v := range calendars {
		i.r.DB.Create(&master.CalendarMaster{
			Base: master.Base{
				ID:      v.ID,
				SameAs:  v.OwlSameAs,
				Context: v.Context,
				Type:    v.Type,
				Date:    parseDate(v.DcDate),
			},
			Title:           v.DcTitle,
			CalendarTitleJa: v.OdptCalendarTitle.Ja,
			CalendarTitleEn: v.OdptCalendarTitle.En,
			Duration:        v.OdptDuration,
		})

		for _, day := range v.OdptDay {
			i.r.DB.Create(&master.CalendarMasterDay{
				CalendarMasterID: v.ID,
				Day:              day,
			})
		}
	}
}

func (i Importer) operator() {
	operators, err := i.uc.GetOperators()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Operators]: %v\n", len(operators))
	i.truncate("operator_masters")

	for _, v := range operators {
		i.r.DB.Create(&master.OperatorMaster{
			Base: master.Base{
				ID:      v.ID,
				SameAs:  v.OwlSameAs,
				Context: v.Context,
				Type:    v.Type,
				Date:    parseDate(v.DcDate),
			},
			Title:           v.DcTitle,
			OperatorTitleJa: v.OdptOperatorTitle.Ja,
			OperatorTitleEn: v.OdptOperatorTitle.En,
		})
	}
}
