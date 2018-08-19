package importer

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/interface/registry"
	"github.com/yuki-toida/refodpt/backend/usecase/raw"
)

const timeFormat = "2006-01-02T15:04:05-07:00"

type Importer struct {
	now time.Time
	uc  *raw.UseCase
	db  *gorm.DB
}

func NewImporter(r *registry.Registry) *Importer {
	return &Importer{
		now: time.Now(),
		uc:  raw.NewUseCase(),
		db:  r.Repository.DB,
	}
}

func (i Importer) Run() {
	result := testing.Benchmark(func(b *testing.B) {
		// i.calendar()
		// i.operator()
		// i.passengerSurvey()
		// i.railDirection()
		// i.railway()
		// i.railwayFare()
		// i.station()
		// i.stationTimetable()
		// i.train()
		// i.trainInformation()
		// i.trainTimetable()
		// i.trainType()
	})
	fmt.Printf("%#v\n", result)
}

func (i Importer) truncate(tableName string) {
	i.db.Exec("TRUNCATE TABLE " + tableName)
}

func parseDate(date string) *time.Time {
	t, err := time.Parse(timeFormat, date)
	if err != nil {
		return nil
	}
	return &t
}

func async(s []interface{}, f func(interface{})) {
	ch := make(chan int, config.Config.DB.Pool/2)
	wg := sync.WaitGroup{}
	for _, v := range s {
		ch <- 1
		wg.Add(1)
		go func(v interface{}) {
			f(v)
			<-ch
			wg.Done()
		}(v)
	}
	wg.Wait()
}
