package collector

import (
	"sync"
	"sync/atomic"

	"github.com/uol/mycenae/lib/gorilla"
	pb "github.com/uol/mycenae/lib/proto"
	"github.com/uol/mycenae/lib/utils"

	"go.uber.org/zap"
)

type Job struct {
	rcvMsg gorilla.TSDBpoint
	i      int
	rPts   *RestErrors
	wg     *sync.WaitGroup
	pts    []*pb.TSPoint
	mm     map[string]*pb.Meta
	mtx    *sync.Mutex
}

// Worker represents the worker that executes the job
type Worker struct {
	JobChannel chan Job
	c          *Collector
}

func (c *Collector) RunWorkers(maxWorkers int) {
	// starting n number of workers
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(c)
		worker.Start()
	}
}

func NewWorker(collector *Collector) Worker {
	return Worker{
		JobChannel: collector.jobQueue,
		c:          collector,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			job := <-w.JobChannel
			// register the current worker into the worker queue.
			ks := "invalid"
			if w.c.isKSIDValid(job.rcvMsg.Tags["ksid"]) {
				ks = job.rcvMsg.Tags["ksid"]
			}

			atomic.AddInt64(&w.c.receivedSinceLastProbe, 1)
			statsPoints(ks, "number")

			packet := &pb.TSPoint{}
			m := &pb.Meta{}

			gerr := w.c.makePoint(packet, m, &job.rcvMsg)
			if gerr != nil {
				atomic.AddInt64(&w.c.errorsSinceLastProbe, 1)

				gblog.Error("makePacket", zap.Error(gerr))
				reu := RestErrorUser{
					Datapoint: job.rcvMsg,
					Error:     gerr.Message(),
				}
				job.mtx.Lock()
				job.rPts.Errors = append(job.rPts.Errors, reu)
				job.mtx.Unlock()

				statsPointsError(ks, "number")
				job.wg.Done()
				return
			}

			ksts := utils.KSTS(m.GetKsid(), m.GetTsid())

			job.mtx.Lock()
			job.pts[job.i] = packet
			if _, ok := job.mm[string(ksts)]; !ok {
				job.mm[string(ksts)] = m
			}
			job.mtx.Unlock()

			job.wg.Done()
		}
	}()
}
