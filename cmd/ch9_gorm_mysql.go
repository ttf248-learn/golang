package main

import (
	"fmt"
	"go-base-learning/data_provider/dao"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "root:sErjf&eNh9zv4CSV@tcp(192.168.19.146:5306)/ykcz_trade?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// db.SingularTable(true) // 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`

	db.AutoMigrate(&dao.MessagesIBDropCopy{})

	// 两个自增函数对象
	tmpString := AutoString()
	tmpNumber := AutoNumber()

	for i := 0; i < 1; i++ {
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

	// 查询数据是否存在
	record := dao.MessagesIBDropCopy{}
	db.First(&record)

	var count int64
	count = 0

	bT := time.Now() // 开始时间

	for i := 0; i < 100000; i++ {
		if err := db.Model(&record).Select("F11_ClOrdID").Update("F11_ClOrdID", tmpString()).Error; err != nil {
			panic(err)
		} else {
			count++
		}

		if count%1000 == 0 {
			eT := time.Since(bT).Microseconds() // 从开始到当前所消耗的时间
			fmt.Println("call ", count, " times, avg cost time(us): ", eT, " tps: ", 1000000/(eT/count))
		}
	}

	fmt.Println("create record success")
}
