package importer

import (
	"time"

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

// Run func
func (i Importer) Run() {
	// i.category()
	// i.calendar()
	// i.operator()
	// i.passengerSurvey()
	// i.railDirection()
	// i.railway()
	// i.railwayFare()
	// i.station()
	i.stationTimetable()
}

func parseDate(date string) *time.Time {
	t, err := time.Parse(timeFormat, date)
	if err != nil {
		return nil
	}
	return &t
}
