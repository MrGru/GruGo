package concensus

import (
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, rf := range rafts {
		if string(k) == string(rf.Leader()) {
			state := r.FormValue("next")
			result := rf.Apply([]byte(state), 1*time.Second)
			if result.Error() != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			newState, ok := result.Response().(string)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if newState != state {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("invalid transition"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result.Response().(string)))
			return
		}
	}
}