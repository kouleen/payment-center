package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/payment"
)

type PaymentRechargeSource struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`                                            // 主键ID
	OrderNo    string     `json:"orderNo" gorm:"column:order_no;not null;"`                                                   // 订单号
	PaySource  *int8      `json:"paySource" gorm:"column:pay_source;not null;default:1"`                                      // 支付来源
	SourceNo   string     `json:"sourceNo" gorm:"column:source_no;not null;"`                                                 // 来源单号
	Remark     string     `json:"remark" gorm:"column:remark;not null;default:''"`                                            // 备注
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`                                        // 删除标志
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`                              // 创建人
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`                              // 修改人
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

func (p *PaymentRechargeSource) TableName() string {
	return "payment_recharge_source"
}

func (p *PaymentRechargeSource) ConvertResp() *payment.PaymentRechargeSourceResponse {
	return &payment.PaymentRechargeSourceResponse{
		Id:         p.ID,
		OrderNo:    p.OrderNo,
		PaySource:  p.SourceNo,
		SourceNo:   p.SourceNo,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
