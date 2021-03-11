package main

import (
	"fmt"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }


func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func Withdraw(amount int) (ans int,err error) {
	sum:=Balance()
	if sum<=amount{
		err=fmt.Errorf("TOO POOR!!!!")
		return 0,err
	}else{
		Deposit(-amount)
		return ans,nil
	}
}


func main()  {
	go teller()
	Deposit(100)
	fmt.Println(Balance())
	_,err:=Withdraw(20)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(Balance())
	return
}