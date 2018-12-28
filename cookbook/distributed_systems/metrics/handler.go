package metrics

import (
	"net/http"
	"time"

	metrics "github.com/rcrowley/go-metrics"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	c := metrics.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func TimerHandler(w http.ResponseWriter, r *http.Request) {
	currt := time.Now()
	t := metrics.GetOrRegisterTimer("timerhandler.timer", nil)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	t.UpdateSince(currt)
}
