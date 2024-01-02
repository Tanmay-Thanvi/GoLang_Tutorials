package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study!")

	presentTime := time.Now()
	fmt.Println("Time : ",presentTime)
	// op : Time :  2023-10-15 12:45:05.783787 +0530 IST m=+0.000362561
	// we need to change the format 

	fmt.Println(presentTime.Format("01-02-2006")) // standard syntax for date
	// 2) 01-02-2006 Monday
	// 3) 01-02-2006 15:04:05 Monday

	createdDate := time.Date(2020,time.August,10,23,23,0,0,time.UTC)
	fmt.Println("Date : ",createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
