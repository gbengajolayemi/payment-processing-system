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
	AddLoan(principal int, time int)
}

// Constants for tax rate and loan calculations
const (
	TaxRate       = 0.05 // Example tax rate of 5%
	BorrowingRate = 50   // Example borrowing rate
	BorrowingOne  = 1    // Example adjustment constant
)

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

// Method to add a loan amount to the account balance and calculate interest
func (a *Account) AddLoan(principal int, time int) {
	if principal > 0 {
		a.Balance += float64(principal)
		interest := loanInterest(principal, time)
		totalPayback := float64(principal + interest)
		fmt.Printf("You have borrowed %d. You will pay back a total of %.2f in interest.\n", principal, totalPayback)
		fmt.Printf("Your new balance is %.2f.\n", a.Balance)
	} else {
		fmt.Println("You have inputted an invalid loan amount.")
	}
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

// ContinuePrompt prompts the user to continue or exit
func ContinuePrompt() bool {
	var response string
	fmt.Print("Do you want to continue? (yes/no): ")
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println("Error reading response. Assuming 'no'.")
		return false
	}
	return response == "yes" || response == "y"
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
		AccountNumber: 1,
		Owner:         "Users",
		Balance:       2000,
	}

	// Interface reference
	var ops AccountOperations = a

	for {
		fmt.Print("Enter your account number: ")
		var inputNumber int
		_, err = fmt.Scanln(&inputNumber)
		if err != nil {
			logger.Printf("Error reading account number: %v", err)
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if inputNumber != a.AccountNumber {
			fmt.Println("Account not found in our system.")
			logger.Printf("Account number %d not found", inputNumber)
			continue
		}

		for {
			fmt.Print("Select option 1 for deposit, 2 for withdraw, 3 for filling tax, 4  loan interest checking, 5 for loan borrowing, 6 for payment processing, 7 for checking balance, or 0 to exit: ")
			var choice int
			_, err = fmt.Scanln(&choice)
			if err != nil {
				logger.Printf("Error reading option: %v", err)
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			switch choice {
			case 1:
				fmt.Print("Enter the amount you want to deposit: ")
				var depositAmount float64
				_, err = fmt.Scanln(&depositAmount)
				if err != nil {
					logger.Printf("Error reading deposit amount: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}
				ops.Deposit(depositAmount)
				logger.Printf("Deposit of %.2f made. New balance: %.2f", depositAmount, a.Balance)

			case 2:
				fmt.Print("Enter the amount you want to withdraw: ")
				var withdrawAmount float64
				_, err = fmt.Scanln(&withdrawAmount)
				if err != nil {
					logger.Printf("Error reading withdrawal amount: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}
				ops.Withdraw(withdrawAmount)
				logger.Printf("Withdrawal of %.2f made. New balance: %.2f", withdrawAmount, a.Balance)

			case 3:
				fmt.Print("Enter the initial amount: ")
				var initialAmount float64
				_, err = fmt.Scan(&initialAmount)
				if err != nil {
					logger.Printf("Error reading initial amount: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				fmt.Print("Enter the interest rate: ")
				var interestRate float64
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
				var principal int
				_, err = fmt.Scan(&principal)
				if err != nil {
					logger.Printf("Error reading principal amount: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				fmt.Print("Enter the number of years you are borrowing for: ")
				var time int
				_, err = fmt.Scan(&time)
				if err != nil {
					logger.Printf("Error reading time: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				result := loanInterest(principal, time)
				fmt.Printf("if you borrow %d you will pay a total of %d in interest in %d year/s.\n", principal, result, time)

			case 5:
				fmt.Print("Enter loan amount: ")
				var loanAmount int
				_, err = fmt.Scanln(&loanAmount)
				if err != nil {
					logger.Printf("Error reading loan amount: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				fmt.Print("Enter the number of years you are borrowing for: ")
				var loanTime int
				_, err = fmt.Scanln(&loanTime)
				if err != nil {
					logger.Printf("Error reading loan time: %v", err)
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				ops.AddLoan(loanAmount, loanTime)
				interest := loanInterest(loanAmount, loanTime)
				totalPayback := float64(loanAmount + interest)
				fmt.Printf("You have borrowed %d. You will pay back a total of %.2f in interest.\n", loanAmount, totalPayback)
				fmt.Printf("Your new balance is %.2f.\n", a.Balance)
				logger.Printf("Loan of %d added. New balance: %.2f", loanAmount, a.Balance)

			case 6:
				fmt.Print("Select payment method: 1 for Credit Card, 2 for Bank Transfer: ")
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
					var cardNumber string
					_, err = fmt.Scanln(&cardNumber)
					if err != nil {
						logger.Printf("Error reading credit card number: %v", err)
						fmt.Println("Invalid input. Please try again.")
						continue
					}
					cc := CreditCardProcessor{CardNumber: cardNumber}
					fmt.Print("Enter amount to process with credit card: ")
					var cardAmount float64
					_, err = fmt.Scanln(&cardAmount)
					if err != nil {
						logger.Printf("Error reading credit card amount: %v", err)
						fmt.Println("Invalid input. Please try again.")
						continue
					}
					ProcessPayment(cc, cardAmount)

				case 2:
					fmt.Print("Enter bank account number: ")
					var bankAccount string
					_, err = fmt.Scanln(&bankAccount)
					if err != nil {
						logger.Printf("Error reading bank account number: %v", err)
						fmt.Println("Invalid input. Please try again.")
						continue
					}
					btp := BankTransferProcessor{BankAccount: bankAccount}
					fmt.Print("Enter amount to process with bank transfer: ")
					var bankAmount float64
					_, err = fmt.Scanln(&bankAmount)
					if err != nil {
						logger.Printf("Error reading bank transfer amount: %v", err)
						fmt.Println("Invalid input. Please try again.")
						continue
					}
					ProcessPayment(btp, bankAmount)

				default:
					fmt.Println("Invalid payment method.")
					continue
				}

			case 7:
				ops.CheckBalance()

			case 0:
				fmt.Println("Exiting...")
				return

			default:
				fmt.Println("Invalid choice. Please try again.")
			}

			// Prompt to continue or exit
			if !ContinuePrompt() {
				fmt.Println("Exiting...")
				return
			}
		}
	}
}
