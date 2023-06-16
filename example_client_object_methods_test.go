package sfdc_test

import (
	"fmt"

	"github.com/stellaraf/go-sfdc"
)

// See client example for client initialization example.
var client *sfdc.Client

func ExampleClient_Account() {
	account, err := client.Account("006G00000678LKJIHE")
	if err != nil {
		// handle error
	}
	fmt.Println(account.ID)
	// Example Output: 006G00000678LKJIHE
}

func ExampleClient_User() {
	user, err := client.Account("003G00000789VUTSRQ")
	if err != nil {
		// handle error
	}
	fmt.Println(user.ID)
	// Example Output: 003G00000789VUTSRQ
}

func ExampleClient_Group() {
	group, err := client.Account("00QG00000890ZYXWVU")
	if err != nil {
		// handle error
	}
	fmt.Println(group.ID)
	// Example Output: 00QG00000890ZYXWVU
}

func ExampleClient_Case() {
	_case, err := client.Case("001G00000987KJIHGF")
	if err != nil {
		// handle error
	}
	fmt.Println(_case.ID)
	// Example Output: 001G00000987KJIHGF
}

func ExampleClient_Contact() {
	contact, err := client.Contact("003G00000123ZYXWVU")
	if err != nil {
		// handle error
	}
	fmt.Println(contact.ID)
	// Example Output: 003G00000123ZYXWVU
}
