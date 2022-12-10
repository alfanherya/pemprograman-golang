package health

import (
	"net/http"
	"time"

	"github.com/dimiro1/health"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type HealthHandler struct {
}

type SelfChecker struct{}

func NewSSelfChecker() *SelfChecker {
	return &SelfChecker{}
}

func HealthAPI(e *echo.Echo) {
	handler := &HealthHandler{}
	e.Add(
		"GET", "/health", handler.HealthChek,
	)
}

func (r *HealthHandler) HealthChek(c echo.Context) error {
	h := health.NewHandler()
	// h.AddChecker("App", NewSSelfChecker())
	if h.Check().IsDown() {
		logrus.WithFields(logrus.Fields{
			"at": time.Now().Format(time.RFC3339),
		}).Errorf("unhealthy service: %s", h.Check())
		return c.JSON(http.StatusServiceUnavailable, h.Check())
	}
	return c.JSON(http.StatusOK, h.Check())
}
