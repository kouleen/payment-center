package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/payment"
	"github.com/kouleen/payment-center/modle"
	"github.com/kouleen/payment-center/utils"
)

func QueryRechargePage(ctx context.Context, req *payment.PaymentRechargeRequest) (resp []*modle.PaymentRecharge, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.PaymentRecharge{}).Where("is_delete = 0")
	if req.GetOrderNo() != "" {
		query = query.Where("order_no like ?", req.OrderNo+"%")
	}
	if req.GetSubject() != "" {
		query = query.Where("subject like ?", "%"+req.Subject+"%")
	}
	if req.PayType != nil {
		query = query.Where("pay_type = ?", req.PayType)
	}
	if req.GetTradeNo() != "" {
		query = query.Where("trade_no = ?", req.GetTradeNo())
	}
	if req.GetFlowNo() != "" {
		query = query.Where("flow_no = ?", req.GetFlowNo())
	}
	if req.WorkflowStatus != nil {
		query = query.Where("workflow_status =?", req.WorkflowStatus)
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc,id desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryRecharge(ctx context.Context, id int64) (resp *modle.PaymentRecharge, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Take(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}
