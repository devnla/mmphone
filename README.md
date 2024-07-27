# Myanmar Phone Number Validator for Go

`mmphone` is a package for validating Myanmar phone numbers in Go, it supports Ooredoo, Atom, MPT, MyTel, and also supports Burmese aphabets.


## Installation

Install `mmphone` with the following command

```bash
go get github.com/devnla/mmphone
```
    
## Usage/Examples

```go
package main

import (
	"fmt"

	"github.com/devnla/mmphone"
)

func main() {
	phoneChecker := mmphone.NewMyanmarPhone()

	phoneNumbers := []string{
		"09977123456",   // Ooredoo
		"+959781234567", // ATOM
		"959440057616",  // MPT WCDMA
		"0949120059",    // MPT CDMA 450 MHz
		"0937128956",    // MPT CDMA 800 MHz
		"09678123456",   // MyTel
		"၀၉၇၇၅၅၄၁၇၉၄",  // ATOM in Burmese numerals
	}

	for _, phoneNumber := range phoneNumbers {
		fmt.Printf("Phone Number: %s\n", phoneNumber)

		// Sanitize phone number
		sanitizedNumber := phoneChecker.SanitizePhoneNumber(phoneNumber)
		fmt.Printf("Sanitized: %s\n", sanitizedNumber)

		// Validate phone number
		isValid := phoneChecker.IsValidMyanmarPhone(phoneNumber)
		fmt.Printf("Is Valid: %v\n", isValid)

		// Get telecom provider
		provider := phoneChecker.GetTelecomName(phoneNumber)
		fmt.Printf("Provider: %s\n", provider)

		// Get network type
		networkType := phoneChecker.GetPhoneNetworkType(phoneNumber)
		fmt.Printf("Network Type: %s\n\n", networkType)
	}
}
```


## Running Tests

To run tests, run the following command

```bash
go test
```


## Acknowledgements

 - [mm_phonenumber](https://github.com/Melomap/mm_phonenumber)


## License

[MIT](LICENSE)

