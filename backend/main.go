package main

import (
	"fmt"
	"net/http"
	"runtime"
	"testing"

	"github.com/bamzi/jobrunner"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/sync/errgroup"

	"github.com/yuki-toida/refodpt/backend/config"
	"github.com/yuki-toida/refodpt/backend/infrastructure/cache"
	"github.com/yuki-toida/refodpt/backend/infrastructure/db"
	"github.com/yuki-toida/refodpt/backend/infrastructure/job/importer"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	admin "github.com/yuki-toida/refodpt/backend/interface/admin/router"
	app "github.com/yuki-toida/refodpt/backend/interface/app/router"
)

func init() {
	config.Init("./", "./.env")
}

func main() {
	db := db.Connect()
	db.LogMode(true)

	cc := cache.NewCache()
	tr := testing.Benchmark(func(b *testing.B) {
		cc.Init(db)
	})
	fmt.Printf("[Benchmark] %v\n", tr)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[MemStats] Alloc:%v TotalAlloc:%v HeapAlloc:%v HeapSys:%v\n", m.Alloc, m.TotalAlloc, m.HeapAlloc, m.HeapSys)

	repo := repository.NewRepository(db, cc)
	jobrunner.Start()
	jobrunner.Schedule("@every 15m", importer.NewImporter(repo))

	adminServer := &http.Server{
		Addr:    ":" + config.Config.AdminServer.Port,
		Handler: admin.Create(repo),
	}
	appServer := &http.Server{
		Addr:    ":" + config.Config.AppServer.Port,
		Handler: app.Create(repo),
	}

	eg := errgroup.Group{}
	eg.Go(func() error { return adminServer.ListenAndServe() })
	eg.Go(func() error { return appServer.ListenAndServe() })

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
