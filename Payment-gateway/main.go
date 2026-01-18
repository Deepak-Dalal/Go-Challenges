/*
The "Payment Gateway" abstraction
Imagine you are building an e-commerce platform. You want to support multiple payment methods (Credit Card, PayPal, and maybe Crypto later). You don't want to change your main logic every time you add a new payment method.

📋 The Requirements:
The Interface: Define a PaymentProcessor interface. It should have a method: 💳

# ProcessPayment(amount float64) bool

The Implementations: Create two different structs that satisfy this interface:

StripeProcessor: Should print "Processing Stripe payment of $X" and always return true.

PayPalProcessor: Should print "Redirecting to PayPal for $X" and always return true.

The Logic: Write a function called Checkout(p PaymentProcessor, amount float64). 🛒

This function should not care which processor it gets. It just calls p.ProcessPayment(amount).

The Test/Main: In your main() function:

Create a slice of PaymentProcessor containing one of each.

Loop through them and call Checkout for each one with an amount of $99.99.
*/
package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

type StripeProcessor struct{}

func (stripeProcessor *StripeProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing Stripe payment of $%.2f", amount)
	return true;
}

type PayPalProcessor struct{}

func (paypalProcessor *PayPalProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Redirecting to PayPal for $%.2f", amount)
	return true;
}

func Checkout(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount);
}
func main() {
	paymentProcessors := []PaymentProcessor{&StripeProcessor{}, &PayPalProcessor{}}

	for _, paymentProcessor:=range paymentProcessors {
		Checkout(paymentProcessor, 99.99)
		fmt.Println();
	}

}