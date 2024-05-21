![stellar](https://res.cloudinary.com/stellaraf/image/upload/v1604277355/stellar-logo-gradient.svg)

## `go-sfdc` [![Tests](https://img.shields.io/github/actions/workflow/status/stellaraf/go-sfdc/tests.yml?style=for-the-badge)](https://github.com/stellaraf/go-sfdc/actions/workflows/tests.yml) [![Go Reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/github.com/stellaraf/go-sfdc) [![GitHub Tag](https://img.shields.io/github/v/tag/stellaraf/go-sfdc?style=for-the-badge&label=Version)](https://github.com/stellaraf/go-sfdc/tags)


## Testing Environment Variables

The following environment variables must be set (and valid) for tests to run.

| Key                          | Description                                          |
| :--------------------------- | :--------------------------------------------------- |
| `SFDC_CLIENT_ID`             | Connected App OAuth2 Client ID (Consumer Key)        |
| `SFDC_CLIENT_SECRET`         | Connected App OAuth2 Client Secret (Consumer Secret) |
| `SFDC_ENCRYPTION_PASSPHRASE` | AES-256 key for encrypting cache values              |
| `SFDC_AUTH_URL`              | Salesforce Authentication URL                        |
| `SFDC_TEST_DATA`             | JSON string of test data, see below                  |

### Test Data

The following values are used in unit tests to fetch and validate data through the `go-sfdc` client methods.

| Key                     | Description                                                                                  |
| :---------------------- | :------------------------------------------------------------------------------------------- |
| `accountId`             | ID of pre-existing Account object                                                            |
| `userId`                | ID of pre-existing User account                                                              |
| `groupId`               | ID of pre-existing group (Queue)                                                             |
| `accountName`           | Name of Account used in the `accountId` field                                                |
| `contactId`             | ID of pre-existing contact that is associated with the Account used in the `accountId` field |
| `accountCustomFieldKey` | Field name of a custom field on the Account object                                           |
| `caseCustomFieldKey`    | Field name of a custom field on the Case object                                              |
| `userEmail`             | Email address of the User used in the `userId` field                                         |
| `serviceContractId`     | ID of pre-existing Service Contract object                                                   |

#### Example

```json
{
    "accountId": "001A000001abcde123",
    "userId": "005A000001xyz7890",
    "groupId": "00G1A00000abcdef12",
    "accountName": "Acme Corp, Inc.",
    "contactId": "003A000001pqrst456",
    "accountCustomFieldKey": "go_sfdc_Test_Field__c",
    "caseCustomFieldKey": "go_sfdc_Test_Field__c",
    "userEmail": "person@example.com",
    "serviceContractId": "0SCA00000lmnop789"
}
```