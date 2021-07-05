package dao

import (
	"time"

	"github.com/shopspring/decimal"
)

type BaseModel struct {
	ID        uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" type:bigint unsigned"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessagesIBDropCopy struct {
	BaseModel
	F11_ClOrdID           string
	F41_OrigClOrdID       string
	F17_ExecID            string
	F20_ExecTransType     string
	F19_ExecRefID         string
	F150_ExecType         string
	F39_OrdStatus         string
	F103_OrdRejReason     string
	F55_Symbol            string
	F207_SecurityExchange string
	F54_Side              string
	F38_OrderQty          decimal.Decimal `sql:"type:decimal(20,10);"`
	F44_Price             decimal.Decimal `sql:"type:decimal(20,10);"`
	F32_LastShares        decimal.Decimal `sql:"type:decimal(20,10);"`
	F31_LastPx            decimal.Decimal `sql:"type:decimal(20,10);"`
	F151_LeavesQty        decimal.Decimal `sql:"type:decimal(20,10);"`
	F14_CumQty            decimal.Decimal `sql:"type:decimal(20,10);"`
	F6_AvgPx              decimal.Decimal `sql:"type:decimal(20,10);"`
	F60_TransactTime      string
	F58_Text              string
	F49_SenderCompID      string
	F56_TargetCompID      string
}

func (MessagesIBDropCopy) TableName() string {
	return "messages_dropcopy"
}
