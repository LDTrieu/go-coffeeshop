package app

import (
	"go-coffeeshop/cmd/proxy/config"

	mylogger "go-coffeeshop/pkg/logger"
)

type App struct {
	logger  *mylogger.Logger
	cfg     *config.Config
	network string
	address string
}

// type CounterServiceServerImpl struct {
// 	gen.UnimplementedProductServiceServer
// 	logger     *mylogger.Logger
// 	rabbitConn *amqp.Connection
// }

// func (g *CounterServiceServerImpl) GetListorderFulfillment(ctx context.Context,
// 	request *gen.GetListOrderFulfillmentRequest) (
// 	*gen.GetListOrderFulfillmentResponse, error) {
// 		g.logger

// }
