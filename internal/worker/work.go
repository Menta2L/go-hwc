package worker

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	hwc "github.com/menta2l/go-hwc/api/hardware/v1"
	"github.com/menta2l/go-hwc/internal/biz"
)

var ProviderSet = wire.NewSet(NewWorker)

var works []Work

type Worker struct {
	h *log.Helper

	euc    *biz.HardwareUsecase
	client hwc.HardwareClient
}

type WorkFunc func(worker *Worker)

func (w *Worker) registerWork(name string, iworkFunc WorkFunc) {
	works = append(works, Work{
		Name:     name,
		WorkFunc: iworkFunc,
	})
}

func (w *Worker) Run(done chan struct{}) {
	wc := len(works)
	cChan := make(chan struct{}, wc)
	for _, work := range works {
		go func(wer *Worker, w Work) {
			w.WorkFunc(wer)
			cChan <- struct{}{}
		}(w, work)
	}
	for i := 0; i < wc; i++ {
		<-cChan
	}
	done <- struct{}{}
}

type Work struct {
	Name     string
	WorkFunc WorkFunc
}

func NewWorker(c hwc.HardwareClient, logger log.Logger, euc *biz.HardwareUsecase) *Worker {
	worker := Worker{
		h:      log.NewHelper(logger),
		euc:    euc,
		client: c,
	}
	worker.registerWork("hello", HelloWork)
	return &worker
}
