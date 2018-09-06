package importer

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/raw"
	"github.com/yuki-toida/refodpt/backend/usecase/tran"
)

type Importer struct {
	db  *gorm.DB
	ruc *raw.UseCase
	tuc *tran.UseCase
}

func NewImporter(r *repository.Repository) *Importer {
	return &Importer{
		db:  r.DB,
		ruc: raw.NewUseCase(),
		tuc: tran.NewUseCase(r),
	}
}

func (i Importer) Run() {
	i.train()
	i.trainInformation()
	i.tuc.UpdateTranAt(time.Now())
}

func (i Importer) Master() {
	i.calendar()
	i.operator()
	i.passengerSurvey()
	i.railDirection()
	i.railway()
	i.railwayFare()
	i.station()
	i.stationTimetable()
	i.trainTimetable()
	i.trainType()
	i.tuc.UpdateMasterAt(time.Now())
}

func (i Importer) truncate(tableName string) {
	i.db.Exec("TRUNCATE TABLE " + tableName)
}

func parse(date string) *time.Time {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", date)
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
