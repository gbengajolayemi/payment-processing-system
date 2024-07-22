package main

import (
	"fmt"
	"log"
	"os"
)

// Account struct representing a bank account
type Account struct {
	AccountNumber int
	Owner         string
	Balance       float64
}

// Define an interface for account operations
type AccountOperations interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	CheckBalance()
}

// Constants for tax rate and loan calculations
const TaxRate = 0.05    // Example tax rate of 5%
const BorrowingRate = 7 // Example borrowing rate
const BorrowingOne = 1  // Example adjustment constant

// Function to calculate final amount and tax
func calculateFinalAmount(initialAmount, interestRate float64) (finalAmount float64, tax float64) {
	finalAmount = initialAmount + (initialAmount * interestRate)
	tax = finalAmount * TaxRate
	return
}

// Function to calculate loan interest
func loanInterest(principal, time int) int {
	interest := principal * (BorrowingOne + (BorrowingRate * time))
	return interest
}

// Method to deposit amount into the account
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
		fmt.Printf("You just deposited %.2f into your account. Your new balance is %.2f.\n", amount, a.Balance)
	} else {
		fmt.Println("You have inputted an invalid amount.")
	}
}

// Method to withdraw amount from the account
func (a *Account) Withdraw(amount float64) {
	if amount > 0 {
		if amount <= a.Balance {
			a.Balance -= amount
			fmt.Printf("You just withdrew %.2f from your account. Your new balance is %.2f.\n", amount, a.Balance)
		} else {
			fmt.Println("Insufficient funds.")
		}
	} else {
		fmt.Println("You have inputted an invalid amount.")
	}
}

// Method to check the account balance
func (a *Account) CheckBalance() {
	fmt.Printf("Your current balance is %.2f.\n", a.Balance)
}

// PaymentProcessor is an interface that defines common payment operations
type PaymentProcessor interface {
	Authorize(amount float64) bool
	Capture(amount float64) bool
	Refund(amount float64) bool
}

// CreditCardProcessor processes payments via credit card
type CreditCardProcessor struct {
	CardNumber string
}

// Authorize authorizes the payment amount for a credit card
func (c CreditCardProcessor) Authorize(amount float64) bool {
	fmt.Printf("Authorizing $%.2f with credit card %s\n", amount, c.CardNumber)
	return true // Assume authorization is successful
}

// Capture captures the payment amount from a credit card
func (c CreditCardProcessor) Capture(amount float64) bool {
	fmt.Printf("Capturing $%.2f from credit card %s\n", amount, c.CardNumber)
	return true // Assume capture is successful
}

// Refund refunds the payment amount to a credit card
func (c CreditCardProcessor) Refund(amount float64) bool {
	fmt.Printf("Refunding $%.2f to credit card %s\n", amount, c.CardNumber)
	return true // Assume refund is successful
}

// BankTransferProcessor processes payments via bank transfer
type BankTransferProcessor struct {
	BankAccount string
}

// Authorize authorizes the payment amount for a bank transfer
func (b BankTransferProcessor) Authorize(amount float64) bool {
	fmt.Printf("Authorizing $%.2f with bank account %s\n", amount, b.BankAccount)
	return true // Assume authorization is successful
}

// Capture captures the payment amount from a bank account
func (b BankTransferProcessor) Capture(amount float64) bool {
	fmt.Printf("Capturing $%.2f from bank account %s\n", amount, b.BankAccount)
	return true // Assume capture is successful
}

// Refund refunds the payment amount to a bank account
func (b BankTransferProcessor) Refund(amount float64) bool {
	fmt.Printf("Refunding $%.2f to bank account %s\n", amount, b.BankAccount)
	return true // Assume refund is successful
}

// ProcessPayment processes the payment using any PaymentProcessor
func ProcessPayment(p PaymentProcessor, amount float64) {
	if p.Authorize(amount) {
		p.Capture(amount)
		fmt.Println("Payment processed successfully.")
	} else {
		fmt.Println("Payment authorization failed.")
	}
}

// Main function
func main() {
	// Create or open a log file
	file, err := os.OpenFile("input.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	// Set up logging to file
	logger := log.New(file, "INFO: ", log.LstdFlags)

	// Create an instance of Account
	a := &Account{
		AccountNumber: 2122,
		Owner:         "Gbenga",
		Balance:       2000,
	}

	// Interface reference
	var ops AccountOperations = a

	var choice int
	var amount float64
	var cardNumber, bankAccount string
	var initialAmount, interestRate float64
	var principal, time int

	for {
		fmt.Print("Enter your account number: ")
		_, err = fmt.Scanln(&choice)
		if err != nil {
			logger.Printf("Error reading account number: %v", err)
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if choice == a.AccountNumber {
			fmt.Println("Account found.")
		} else {
			fmt.Println("Account not found in our system.")
			logger.Printf("Account number %d not found", choice)
			continue
		}

		fmt.Print("Select option 1 for deposit, option 2 for withdraw, option 3 for filling tax, option 4 for loan interest, option 5 for payment processing, 6 for checking balance, or 0 to exit: ")
		_, err = fmt.Scanln(&choice)
		if err != nil {
			logger.Printf("Error reading option: %v", err)
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter the amount you want to deposit: ")
			_, err = fmt.Scanln(&amount)
			if err != nil {
				logger.Printf("Error reading deposit amount: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}
			ops.Deposit(amount)
			logger.Printf("Deposit of %.2f made. New balance: %.2f", amount, a.Balance)

		case 2:
			fmt.Print("Enter the amount you want to withdraw: ")
			_, err = fmt.Scanln(&amount)
			if err != nil {
				logger.Printf("Error reading withdrawal amount: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}
			ops.Withdraw(amount)
			logger.Printf("Withdrawal of %.2f made. New balance: %.2f", amount, a.Balance)

		case 3:
			fmt.Print("Enter the initial amount: ")
			_, err = fmt.Scan(&initialAmount)
			if err != nil {
				logger.Printf("Error reading initial amount: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			fmt.Print("Enter the interest rate: ")
			_, err = fmt.Scan(&interestRate)
			if err != nil {
				logger.Printf("Error reading interest rate: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			finalAmount, tax := calculateFinalAmount(initialAmount, interestRate)
			fmt.Printf("The final amount is %.2f and the tax is %.2f\n", finalAmount, tax)

		case 4:
			fmt.Print("Enter the amount you want to borrow: ")
			_, err = fmt.Scan(&principal)
			if err != nil {
				logger.Printf("Error reading principal amount: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			fmt.Print("Enter the number of years you are borrowing for: ")
			_, err = fmt.Scan(&time)
			if err != nil {
				logger.Printf("Error reading time period: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			result := loanInterest(principal, time)
			fmt.Printf("If you borrow %d, you will pay back %d in %d years.\n", principal, result, time)

		case 5:
			fmt.Print("Select payment method (1 for Credit Card, 2 for Bank Transfer): ")
			var paymentMethod int
			_, err = fmt.Scanln(&paymentMethod)
			if err != nil {
				logger.Printf("Error reading payment method: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			switch paymentMethod {
			case 1:
				fmt.Print("Enter credit card number: ")
				fmt.Scanln(&cardNumber)
				cc := CreditCardProcessor{CardNumber: cardNumber}
				fmt.Print("Enter amount to process with credit card: ")
				fmt.Scanln(&amount)
				ProcessPayment(cc, amount)
				logger.Printf("Processed payment with credit card for amount: %.2f", amount)

			case 2:
				fmt.Print("Enter bank account number: ")
				fmt.Scanln(&bankAccount)
				bank := BankTransferProcessor{BankAccount: bankAccount}
				fmt.Print("Enter amount to process with bank transfer: ")
				fmt.Scanln(&amount)
				ProcessPayment(bank, amount)
				logger.Printf("Processed payment with bank transfer for amount: %.2f")

			default:
				fmt.Println("Invalid payment method selected.")
				logger.Printf("Invalid payment method: %d", paymentMethod)
			}

		case 6:
			ops.CheckBalance()

		case 0:
			fmt.Println("Exiting application.")
			logger.Println("Application exited.")
			return

		default:
			fmt.Println("Invalid option selected.")
			logger.Printf("Invalid option selected: %d", choice)
		}

		fmt.Print("Do you want to process another transaction? (yes/no): ")
		var continueChoice string
		_, err = fmt.Scanln(&continueChoice)
		if err != nil || (continueChoice != "yes" && continueChoice != "no") {
			logger.Printf("Error or invalid input for continue choice: %v", err)
			fmt.Println("Invalid input. Exiting application.")
			return
		}

		if continueChoice == "no" {
			fmt.Println("Exiting application.")
			logger.Println("Application exited.")
			return
		}
	}
}
