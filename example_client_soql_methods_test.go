package sfdc_test

import (
	"fmt"
)

func ExampleClient_OpenCases() {
	openCases, err := client.OpenCases()
	if err != nil {
		// handle error
	}
	for _, openCase := range openCases {
		fmt.Printf("Id=%s, Subject=%s, OwnerId=%s", openCase.ID, openCase.Subject, openCase.OwnerID)
	}
	// Output:
	// Id=003G000002z8VBQIA2, Subject=Case Subject, OwnerId=00QG00000234ZYXWVU
}

func ExampleClient_UserName() {
	name, err := client.UserName("001G00000123ABCDEF")
	if err != nil {
		// handle error
	}
	fmt.Println(name)
	// Output: John Doe
}

func ExampleClient_GroupName() {
	name, err := client.GroupName("006G00000345LKJIHG")
	if err != nil {
		// handle error
	}
	fmt.Println(name)
	// Output: Help Desk
}

func ExampleClient_AccountIDFromName() {
	id, err := client.AccountIDFromName("Acme Corp")
	if err != nil {
		// handle error
	}
	fmt.Println(id)
	// Output: 001G00000456MNOPQR
}

func ExampleClient_Customers() {
	customers, err := client.Customers()
	if err != nil {
		// handle error
	}
	for _, customer := range customers {
		fmt.Printf("Id=%s, Name=%s, Type=%s, ServiceIdentifier=%s", customer.ID, customer.Name, customer.Type, customer.ServiceIdentifier)
		// Output: Id=00QG00000567DCBAZY, Name=Acme Corp, Type=Customer, ServiceIdentifier=0123456
		// ...
	}
}
