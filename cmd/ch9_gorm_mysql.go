package main

import (
	"fmt"
	"go-base-learning/data_provider/dao"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

func AutoString() func() string {
	i := 0
	return func() string {
		i++
		return strconv.Itoa(i)
	}
}

func AutoNumber() func() decimal.Decimal {
	var i int32
	i = 0
	return func() decimal.Decimal {
		i++
		return decimal.NewFromInt32(i)
	}
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`

	db.AutoMigrate(&dao.MessagesIBDropCopy{})

	tmpString := AutoString()
	tmpNumber := AutoNumber()
	for i := 0; i < 10000; i++ {
		idxMessageIBDropCopy := dao.MessagesIBDropCopy{
			F11_ClOrdID:           tmpString(),
			F41_OrigClOrdID:       tmpString(),
			F17_ExecID:            tmpString(),
			F20_ExecTransType:     tmpString(),
			F19_ExecRefID:         tmpString(),
			F150_ExecType:         tmpString(),
			F39_OrdStatus:         tmpString(),
			F103_OrdRejReason:     tmpString(),
			F55_Symbol:            tmpString(),
			F207_SecurityExchange: tmpString(),
			F54_Side:              tmpString(),
			F38_OrderQty:          tmpNumber(),
			F44_Price:             tmpNumber(),
			F32_LastShares:        tmpNumber(),
			F31_LastPx:            tmpNumber(),
			F151_LeavesQty:        tmpNumber(),
			F14_CumQty:            tmpNumber(),
			F6_AvgPx:              tmpNumber(),
			F60_TransactTime:      tmpString(),
			F58_Text:              tmpString(),
			F49_SenderCompID:      tmpString(),
			F56_TargetCompID:      tmpString(),
		}

		if err := db.Create(&idxMessageIBDropCopy).Error; err != nil {
			panic(err)
		}
	}

	fmt.Println("create record success")
}
