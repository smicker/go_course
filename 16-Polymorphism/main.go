package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type CameraSales struct {
	cameraType   string
	pricePerUnit int
	salesAmount  int
}

func (cs CameraSales) calculate() int {
	return cs.pricePerUnit * cs.salesAmount
}
func (cb CameraSales) source() string {
	return cb.cameraType
}

type AppSales struct {
	appName string
	profit  int
}

func (as AppSales) calculate() int {
	return as.profit
}
func (as AppSales) source() string {
	return as.appName
}

func main() {
	project1 := CameraSales{cameraType: "Wallet", pricePerUnit: 9000, salesAmount: 1000}
	project2 := AppSales{appName: "DID", profit: 2500000}
	incomes := []Income{project1, project2}
	calculateNetIncome(incomes)
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = %d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = %d\n", netincome)
}
