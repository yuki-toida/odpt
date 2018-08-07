package importer

import (
	"fmt"
	"time"

	"github.com/yuki-toida/refodpt/backend/domain/model"
	"github.com/yuki-toida/refodpt/backend/domain/repository"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
	"github.com/yuki-toida/refodpt/backend/usecase/raw"
)

const timeFormat = "2006-01-02T15:04:05-07:00"

// Importer struct
type Importer struct {
	now time.Time
	uc  *raw.UseCase
	r   repository.Repository
}

// NewImporter func
func NewImporter(r *registry.Registry) *Importer {
	return &Importer{
		now: time.Now(),
		uc:  raw.NewUseCase(),
		r:   r.Repository,
	}
}

func parseDate(date string) *time.Time {
	var result *time.Time
	if t, err := time.Parse(timeFormat, date); err == nil {
		result = &t
	}
	return result
}

// Run func
func (i Importer) Run() {
	i.calendar()
	i.operator()
}

func (i Importer) calendar() {
	i.r.Delete(model.Calendar{})
	i.r.Delete(model.CalendarDay{})

	calendars, err := i.uc.GetCalendar()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Calendars]: %v\n", len(calendars))

	for _, v := range calendars {
		i.r.Create(model.Calendar{
			ID:              v.ID,
			SameAs:          v.OwlSameAs,
			Context:         v.Context,
			Type:            v.Type,
			Date:            parseDate(v.DcDate),
			Title:           v.DcTitle,
			CalendarTitleJa: v.OdptCalendarTitle.Ja,
			CalendarTitleEn: v.OdptCalendarTitle.En,
			Duration:        v.OdptDuration,
			UpdatedAt:       i.now,
		})

		for _, day := range v.OdptDay {
			i.r.Create(model.CalendarDay{
				CalendarID: v.ID,
				Day:        day,
				UpdatedAt:  i.now,
			})
		}
	}
}

func (i Importer) operator() {
	i.r.Delete(model.Operator{})

	operators, err := i.uc.GetOperator()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Operators]: %v\n", len(operators))

	for _, v := range operators {
		i.r.Create(model.Operator{
			ID:              v.ID,
			SameAs:          v.OwlSameAs,
			Context:         v.Context,
			Type:            v.Type,
			Date:            parseDate(v.DcDate),
			Title:           v.DcTitle,
			OperatorTitleJa: v.OdptOperatorTitle.Ja,
			OperatorTitleEn: v.OdptOperatorTitle.En,
			UpdatedAt:       i.now,
		})
	}
}
