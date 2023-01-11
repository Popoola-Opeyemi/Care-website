package routes

import (
	"github.com/Popoola-Opeyemi/sprintBackend/models"
	"github.com/Popoola-Opeyemi/sprintBackend/utils"

	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// AccessPass describes the minimum role required
// for a request to be fielded.
// for MinRole i, any role j > i will not succeed
// for a request with a http method m ∈ MethodSet
type AccessPass struct {
	MinRole   int64
	MethodSet map[string]bool
}

// CreatePass create a constraint for a request
func CreatePass(minRole int64, methods ...string) AccessPass {
	M := make(map[string]bool)
	for _, m := range methods {
		M[m] = true
	}
	return AccessPass{minRole, M}
}

// GatePass ensures the constraint described in AccessPass is met
// for each request and returns a 401 unauthorized status code if violated
// and logs the status of the request
func GatePass(pass AccessPass, log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			log.Info(req.Method, zap.String("path", req.URL.Path))
			role, err := utils.GetRoleNumFromEchoSession(c, models.AuthSessionName)
			if err != nil {
				log.Debug(
					"getting role from session",
					zap.String("error", err.Error()),
				)
				return err
			}
			// if the current request method exists in pass, check if
			// request should be completed using constraint on MinRole
			if pass.MethodSet[req.Method] {
				if role > pass.MinRole {
					log.Debug(
						"access denied",
						zap.Int64("role", role),
						zap.String("path", req.URL.Path),
						zap.String("method", req.Method),
					)
					return c.String(http.StatusUnauthorized, "access denied")
				}
			}
			log.Info(
				"access granted",
				zap.String("path", req.URL.Path),
				zap.String("method", req.Method),
			)
			return next(c)
		}
	}
}
