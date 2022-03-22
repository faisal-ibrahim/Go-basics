package main

func main() {
	println("If Statements in Go.....")

	var customerHeight int = 140
	customerAge := 18

	if customerHeight >= 150 || customerAge >= 18 {
		println("Can access ride")
	} else if customerHeight >= 120 {
		println("Customer can access children's rides")
	} else {
		println("Customer is too small")
	}
}
