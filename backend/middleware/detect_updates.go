package middleware

import (
	"backend/services"

	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Need a better name for the function
func DetectQuestionUpdates(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Request().Method {
		case http.MethodPost:
		case http.MethodPut:
		case http.MethodPatch:
		case http.MethodDelete:
		default:
			return next(c)
		}

		c.Response().Before(func() {
			// TODO: This seems to pass even if the final status is not >=200 && < 300
			// I think it might be because we're using Response#Before instead
			// of Response#After. Will look into it later as it's not really a
			// problem, it will just send updates more often than it should.
			if c.Response().Status/100 == 2 {
				services.NotifyUpdateListeners()
			}
		})
		return next(c)
	}
}
