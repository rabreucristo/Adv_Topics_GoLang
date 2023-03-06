package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var numberOfDepositsChequing int

var numberOfWithdrawalsChequing int

var numberOfDepositsSavings int

var numberOfWithdrawalsSavings int

var interest float32 = 10

var fee float32 = 5

// Chequing Account "instantiation" and methods

type Chequing struct {
	balance      float32
	fundsMessage string
}

func (c Chequing) getBalanceChequing() float32 {

	return c.balance
}

func (c *Chequing) deposit(amount float32) float32 {

	numberOfDepositsChequing++

	c.balance += amount

	c.fundsMessage = ""

	return c.balance

}

func (c *Chequing) withdrawal(amount float32) (float32, error) {

	numberOfWithdrawalsChequing++

	if amount > c.balance {
		c.fundsMessage = "No Funds Available"
		return c.balance, errors.New("Insufficient Funds")
	} else {
		if numberOfWithdrawalsChequing%5 == 0 && numberOfWithdrawalsChequing > 0 {
			if amount+fee > c.balance {
				c.fundsMessage = "No Funds Available"
				return 0, errors.New("Insufficient Funds")
			} else {
				c.balance -= (amount + fee)
			}
		} else {
			c.balance -= amount
		}
	}
	c.fundsMessage = ""
	return c.balance, nil

}

func (c Chequing) printBalance() string {
	var balance = c.getBalanceChequing()

	if c.fundsMessage != "" {
		return fmt.Sprintf(c.fundsMessage)
	} else {
		var msgBalance = ""

		if numberOfWithdrawalsChequing%5 == 0 && numberOfWithdrawalsChequing > 0 {
			msgBalance = fmt.Sprintf("Your balance is $%.2f and an interest fee of $%.2f", balance, fee)
		} else {
			msgBalance = fmt.Sprintf("Your balance is $%.2f ", balance)
		}
		return msgBalance
	}

}

// End of Chequing account

// Savings Account "Instantiation" and methods

type Savings struct {
	balance float32
}

func (s Savings) getBalanceSavings() float32 {

	return s.balance
}

func (s *Savings) deposit(amount float32) float32 {

	numberOfDepositsSavings++

	if numberOfDepositsSavings%5 == 0 && numberOfDepositsSavings > 0 {
		s.balance += amount + interest
	} else {
		s.balance += amount
	}

	return s.balance

}

func (s *Savings) withdrawal(amount float32) (float32, error) {

	numberOfWithdrawalsSavings++

	if amount > s.balance {
		return 0, errors.New("Insufficient Funds")
	} else {

		s.balance -= amount

	}
	return s.balance, nil

}

func (s Savings) printBalance() string {
	var balance = s.getBalanceSavings()

	var msgBalance = ""

	if numberOfDepositsSavings%5 == 0 && numberOfDepositsSavings > 0 {
		msgBalance = fmt.Sprintf("Your balance is $%.2f and an interest of $%.2f", balance, interest)
	} else {
		msgBalance = fmt.Sprintf("Your balance is $%.2f ", balance)
	}
	return msgBalance
}

func main() {

	fmt.Println("Welcome to the International Banking Transfer!!")
	c := Chequing{balance: 100}

	s := Savings{balance: 100}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Please Choose your account type: 1-> Chequing 2-> Savings 3-> exit")

		scanner.Scan()

		account, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Provide an integer number", err)
		} else {
			switch account {

			case 1:

				fmt.Println("Would like to make a  1->Deposit or a 2->Withdrawal? 3->Exit")

				scanner.Scan()

				action, err := strconv.Atoi(scanner.Text())

				if err != nil {
					fmt.Println("Provide an integer number", err)

				} else {

					fmt.Println("Enter the amount: ")

					scanner.Scan()

					amount, err := strconv.ParseFloat(scanner.Text(), 32)

					if err != nil {
						fmt.Println("Provide a number", err)
					} else {

						if action == 1 {

							c.deposit(float32(amount))

							fmt.Println(c.printBalance())

						} else if action == 2 {
							c.withdrawal(float32(amount))
							fmt.Println(c.printBalance())
						}

					}

				}

			case 2:

				fmt.Println("Would like to make a  1->Deposit or a 2->Withdrawal? 3->Exit")

				scanner.Scan()

				action, err := strconv.Atoi(scanner.Text())

				if err != nil {
					fmt.Println("Provide an integer number", err)

				} else {

					fmt.Println("Enter the amount: ")

					scanner.Scan()

					amount, err := strconv.ParseFloat(scanner.Text(), 32)

					if err != nil {
						fmt.Println("Provide a number", err)
					} else {

						if action == 1 {

							s.deposit(float32(amount))

							fmt.Println(s.printBalance())

						} else if action == 2 {
							s.withdrawal(float32(amount))
							fmt.Println(s.printBalance())
						}

					}

				}

			case 3:
				fmt.Println("Thanks for banking with us! Have a great Day!")
				os.Exit(0)
			default:
				fmt.Println("Choose an option between 1 and 3")
				continue
			}
		}

	}

}
