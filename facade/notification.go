package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}
func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Seding wallet debit notification")
}
