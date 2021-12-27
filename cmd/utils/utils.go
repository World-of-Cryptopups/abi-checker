package utils

import (
	"regexp"

	"github.com/World-of-Cryptopups/abi-checker/cmd/internal"
)

// Gets the struct types of `n`
func FilterStructTypes(structs []internal.ABIStructProps, n string) internal.ABIStructProps {
	for _, v := range structs {
		if v.Name == n {
			return v
		}
	}

	return internal.ABIStructProps{}
}

// IsValidBase32 checks if `v` is a valid base32 string.
func IsValidBase32(v string) bool {
	// https://github.com/EOSIO/eos/issues/955#issuecomment-351866599
	r := regexp.MustCompile(`^[a-z][a-z1-5\.]{0,10}([a-z1-5]|^\.)[a-j1-5]?$`)

	return r.MatchString(v)
}
