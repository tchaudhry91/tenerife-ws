package internal

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// WrapLogger wraps an internal Controller with a logger
func WrapLogger(logger *logrus.Logger, controller http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			logger.Info("Route:", r.URL.Path, "\tTook", time.Since(begin))
		}(time.Now())
		controller(w, r)
	}
}
