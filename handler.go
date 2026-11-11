package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/payment"
	"github.com/kouleen/payment-center/handle"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// QueryPaymentRechargePage implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) QueryPaymentRechargePage(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargePageResponse, err error) {
	return handle.QueryRechargePage(ctx, req)
}

// QueryPaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) QueryPaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargeResponse, err error) {
	return handle.QueryRecharge(ctx, req)
}

// SyncPaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) SyncPaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp *payment.PaymentRechargeResponse, err error) {
	return
}

// CreatePaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CreatePaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp bool, err error) {
	return
}

// UpdatePaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) UpdatePaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp bool, err error) {
	return
}

// CommitPaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CommitPaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp bool, err error) {
	return
}

// ClosePaymentRecharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) ClosePaymentRecharge(ctx context.Context, req *payment.PaymentRechargeRequest) (resp bool, err error) {
	return
}
