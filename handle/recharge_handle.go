package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/payment"
	"github.com/kouleen/payment-center/service"
)

func QueryRechargePage(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargePageResponse, err error) {
	return service.QueryRechargePage(ctx, req)
}

func QueryRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargeResponse, err error) {
	return service.QueryRecharge(ctx, req)
}
