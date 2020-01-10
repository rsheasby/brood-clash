package controllers

import (
	"backend/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func init() {
	ControllerConfigList.Register(ConfigureAuthTestAPI)
}

func ConfigureAuthTestAPI(api *operations.BroodClashAPI) {
	api.AuthTestHandler = operations.AuthTestHandlerFunc(func(
		params operations.AuthTestParams,
		principal interface{},
	) middleware.Responder {
		return operations.NewAuthTestNoContent()
	})
}
