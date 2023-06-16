package sfdc_test

import "fmt"

func ExampleClient_AccountContact() {
	contact, err := client.AccountContact("00QG00000234TSRQPO")
	if err != nil {
		// handle error
	}
	fmt.Println(contact.ID)
	// Example Output: 006G00000345ONMLKJ
}
