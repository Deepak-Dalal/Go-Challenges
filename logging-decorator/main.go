/*
The "Logging Decorator" 📝
We are going to take your PaymentProcessor from the previous round and add a "Logging Layer" to it without touching the code inside StripeProcessor or PayPalProcessor.

📋 The Requirements:
The Wrapper: Create a new struct called LoggingMiddleware. 📦

This struct should have a field that holds a PaymentProcessor (the interface).

The Implementation: Make LoggingMiddleware satisfy the PaymentProcessor interface as well.

In its ProcessPayment method, it should:

Print: "--- LOG: Starting payment for $X ---" 🪵

Call the inner processor's ProcessPayment method and save the result.

Print: "--- LOG: Payment completed. Result: [true/false] ---" ✅

The Wiring: In your main() function:

Wrap your StripeProcessor inside a LoggingMiddleware.

Pass that wrapped version into the Checkout function.
*/
package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

type StripeProcessor struct{}

func (stripeProcessor *StripeProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	return true;
}

type PayPalProcessor struct{}

func (paypalProcessor *PayPalProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Redirecting to PayPal for $%.2f\n", amount)
	return true;
}

func Checkout(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount);
}

type LoggingMiddleware struct{
	Inner PaymentProcessor
}

func (loggingMiddleware *LoggingMiddleware) ProcessPayment(amount float64) bool {
	fmt.Printf("--- LOG: Starting payment for $%.2f ---\n", amount )
	result:= loggingMiddleware.Inner.ProcessPayment(amount);
	fmt.Printf("--- LOG: Payment completed. Result: %t ---", result )
	return result;
}

func main() {
	loggingMiddleware :=&LoggingMiddleware{Inner: &StripeProcessor{}}
	Checkout(loggingMiddleware, 99.99)
}