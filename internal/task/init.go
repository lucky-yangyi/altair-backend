package task

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func CronInit() {
	log.Println("Starting...")

	c := cron.New()
	//定时归集钱包余额
	err := c.AddFunc("1 */30 * * * *", func() {
		CollectTransactionTask()
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	//getBalance
	err = c.AddFunc("1 * * * * *", func() {
		getBalance()
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	//c.AddFunc("0 */1 * * * ?",service.GetMonthBillService().AddMonthBill)
	c.Start()
}
