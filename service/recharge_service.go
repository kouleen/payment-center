package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/payment"
	"github.com/kouleen/payment-center/repository"
)

func QueryRechargePage(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargePageResponse, err error) {
	list, total, err := repository.QueryRechargePage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*payment.PaymentRechargeResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &payment.PaymentRechargePageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargeResponse, err error) {
	item, err := repository.QueryRecharge(ctx, req.GetId())
	if err != nil {
		return
	}
	orderNo := item.OrderNo
	itemSource, err := repository.QueryRechargeSourceByOrderNo(ctx, orderNo)
	if err != nil {
		return
	}
	resp = item.ConvertResp()
	resp.RechargeSource = itemSource.ConvertResp()
	return resp, nil
}
