package main

import "fmt"

type Payment interface {
	Pay()
}

func ProccessPayment(p Payment) {
	p.Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

/*Soluciòn*/
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() { //Adaptacion
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProccessPayment(cash)

	//bank no puede implementar correctamente Pay
	//bank := &BankPayment{}
	//ProccessPayment(bank.)
	// al ejecutar esto arroja un error

	// *****************************
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProccessPayment(bpa)
}
