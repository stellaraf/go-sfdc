// Methods that consume object _and_ SOQL methods, or that handle data processing.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
)

// Retrieve the primary contact for an account.
func (client *Client) AccountContact(accountId string) (contact types.Contact, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	account, err := client.Account(accountId)
	if err != nil {
		return
	}
	if account.MainPointOfContact == "" {
		err = fmt.Errorf("account '%s' (%s) has no Main Point of Contact associated", accountId, account.Name)
		return
	}
	contact, err = client.Contact(account.MainPointOfContact)
	return
}
