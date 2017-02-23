package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iain17/dht-hello/service/restapi/operations"
	"github.com/iain17/dht-hello/service/models"
	"github.com/go-openapi/swag"
	logger "github.com/Sirupsen/logrus"
)

func StartSearch(params operations.StartSearchParams) middleware.Responder {
	err := dht.NewSearch(params.Identifier, *params.Port, *params.ImpliedPort)
	if err != nil {
		logger.Debug(err)
		return operations.NewStartSearchDefault(500).WithPayload(&models.Error{
			Message: swag.String(err.Error()),
		})
	}
	return operations.NewStartSearchOK()
}