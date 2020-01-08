package middleware

import (
	"backend/restapi/operations"
	"backend/services"
	"github.com/go-openapi/errors"
)

func ConfigureAuth(api *operations.BroodClashAPI) {
	api.APIKeyAuth = func(token string) (interface{}, error) {
		if services.ValidateLoginCode(token) {
			return true, nil
		} else {
			return nil, errors.Unauthenticated("Invalid login code")
		}
	}
}
