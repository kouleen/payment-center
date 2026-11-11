package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/payment-center/modle"
)

func QueryRechargeSourceByOrderNo(ctx context.Context, orderNo string) (resp *modle.PaymentRechargeSource, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Take(&resp, orderNo).Error; err != nil {
		return nil, err
	}
	return
}
