package importer

import (
	"fmt"

	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/domain/model/raw"
)

func (i Importer) calendar() {
	calendars, err := i.uc.GetCalendars()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Calendars]: %v\n", len(calendars))
	i.truncate("calendar_masters")
	i.truncate("calendar_master_days")

	s := make([]interface{}, len(calendars))
	for i, v := range calendars {
		s[i] = v
	}
	async(s, i.createCalendar)
}

func (i Importer) createCalendar(calendar interface{}) {
	v, _ := calendar.(raw.Calendar)
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
