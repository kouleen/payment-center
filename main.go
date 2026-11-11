package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kouleen/common/bootstrap"
	"github.com/kouleen/common/middleware"
	"github.com/kouleen/idl/kitex_gen/payment/paymentservice"
	"github.com/kouleen/idl/kitex_gen/rpc"
)

func main() {
	bootstrap.Run(rpc.PAYMENT_RPC_SERVER, func(option ...server.Option) server.Server {
		return paymentservice.NewServer(new(PaymentServiceImpl), option...)
	}, bootstrap.WithServerMiddleware(middleware.RpcServerMiddleware))
}
