package main

import (
	"encoding/json"
	"fmt"
	"kopever/bankcore/bank"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}

	accounts[1002] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Mark",
				Address: "Irvine, California",
				Phone:   "(949) 555 0198",
			},
			Number: 1002,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			// json.NewEncoder(w).Encode(bank.Statement(account))
			fmt.Fprint(w, account.Statement())
		}
	}
}

// CustomAccount ...
type CustomAccount struct {
	*bank.Account
}

// Statement ...
func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	fromqs := req.URL.Query().Get("from")
	toqs := req.URL.Query().Get("to")
	amountqs := req.URL.Query().Get("amount")

	if fromqs == "" {
		fmt.Fprint(w, "From account number is missing!")
	} else if toqs == "" {
		fmt.Fprint(w, "To account number is missing!")
	} else if amountqs == "" {
		fmt.Fprint(w, "Amount is missing!")
	} else {
		if fromNumber, err := strconv.ParseFloat(fromqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid from account number!")
		} else if toNumber, err := strconv.ParseFloat(toqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid to account number!")
		} else {
			fromAccount, ok := accounts[fromNumber]
			if !ok {
				fmt.Fprintf(w, "From account with number %v can't be found!", fromqs)
			}

			toAccount, ok := accounts[toNumber]
			if !ok {
				fmt.Fprintf(w, "To account with number %v can't be found!", toqs)
			}

			if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
				fmt.Fprintf(w, "Invalid amount number!")
			} else {
				err := fromAccount.Transfer(amount, toAccount.Account)
				if err != nil {
					fmt.Fprint(w, err)
				} else {
					fmt.Fprint(w, fromAccount.Statement(), "\n", toAccount.Statement())
				}
			}
		}
	}
}
