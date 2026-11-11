package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/payment"
)

type PaymentRecharge struct {
	ID             int64      `json:"id,string" gorm:"column:id;primary_key;not null"`                                            // 主键ID
	OrderNo        string     `json:"orderNo" gorm:"column:order_no;not null;uniqueIndex"`                                        // 订单号
	Amount         int64      `json:"amount" gorm:"column:amount;not null"`                                                       // 充值金额 单位(分)
	PayFee         int64      `json:"payFee" gorm:"column:pay_fee;not null"`                                                      // 手续费 单位(分)
	Rate           string     `json:"rate" gorm:"column:rate;not null;default:'0.006'"`                                           // 本次交易费率
	Subject        string     `json:"subject" gorm:"column:subject"`                                                              // 订单标题
	QrCode         string     `json:"qrCode" gorm:"column:qr_code;default:''"`                                                    // 二维码链接/二维码内容(qr_code)
	PayType        *int8      `json:"payType" gorm:"column:pay_type;default:1"`                                                   // 支付类型 1 支付宝; 2 微信;
	PayTime        *time.Time `json:"payTime" gorm:"column:pay_time;"`                                                            // 付款时间
	NotifyTime     *time.Time `json:"notifyTime" gorm:"column:notify_time;"`                                                      // 回调时间
	TradeNo        string     `json:"tradeNo" gorm:"column:trade_no;default:''"`                                                  // 第三方交易号
	FlowNo         string     `json:"flowNo" gorm:"column:flow_no;default:''"`                                                    // 流程单号
	WorkflowStatus *int8      `json:"workflowStatus" gorm:"column:workflow_status;not null;default:0"`                            // 审批状态
	ReviewTime     *time.Time `json:"reviewTime" gorm:"column:review_time;"`                                                      // 审批时间
	Status         *int8      `json:"status" gorm:"column:status;not null;default:1"`                                             // 订单状态
	Remark         string     `json:"remark" gorm:"column:remark;not null;default:''"`                                            // 备注
	IsDelete       *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`                                        // 删除标志
	CreatedBy      int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`                              // 创建人
	UpdatedBy      int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`                              // 修改人
	CreateTime     *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdateTime     *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

func (p *PaymentRecharge) TableName() string {
	return "payment_recharge"
}

func (p *PaymentRecharge) ConvertResp() *payment.PaymentRechargeResponse {
	var payTimeUnixMilli int64
	if p.PayTime != nil {
		payTimeUnixMilli = p.PayTime.UnixMilli()
	}
	var notifyTimeUnixMilli int64
	if p.NotifyTime != nil {
		notifyTimeUnixMilli = p.NotifyTime.UnixMilli()
	}
	var reviewTimeUnixMilli int64
	if p.ReviewTime != nil {
		reviewTimeUnixMilli = p.ReviewTime.UnixMilli()
	}
	return &payment.PaymentRechargeResponse{
		Id:             p.ID,
		OrderNo:        p.OrderNo,
		Amount:         p.Amount,
		PayFee:         p.PayFee,
		Rate:           p.Rate,
		Subject:        p.Subject,
		QrCode:         p.QrCode,
		PayType:        p.PayType,
		PayTime:        payTimeUnixMilli,
		NotifyTime:     notifyTimeUnixMilli,
		TradeNo:        p.TradeNo,
		FlowNo:         p.FlowNo,
		WorkflowStatus: p.WorkflowStatus,
		ReviewTime:     reviewTimeUnixMilli,
		Status:         p.Status,
		Remark:         p.Remark,
		IsDelete:       p.IsDelete,
		CreatedBy:      p.CreatedBy,
		UpdatedBy:      p.UpdatedBy,
		CreateTime:     p.CreateTime.UnixMilli(),
	}
}
