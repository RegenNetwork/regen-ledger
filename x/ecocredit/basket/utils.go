package basket

import (
	"fmt"
	"regexp"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

const (
	nameMinLen  = 3
	nameMaxLen  = 8
	denomPrefix = "eco"
)

var (
	// RegexBasketName requires the first character to be alphabetic, the rest can be alphanumeric. We
	// reduce length constraints by one to account for the first character being forced to alphabetic.
	RegexBasketName = fmt.Sprintf(`[a-zA-Z][a-zA-Z0-9]{%d,%d}`, nameMinLen-1, nameMaxLen-1)
	// RegexBasketDenom requires the first part to match the denom prefix, the second part to match the
	// format for a credit type abbreviation (with or without an exponent prefix), and the third part to
	// satisfy the basket name. Each of the three parts must also be separated by a ".".
	RegexBasketDenom = fmt.Sprintf(`%s.[a-zA-Z]{1,4}.%s`, denomPrefix, RegexBasketName)

	regexBasketName  = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexBasketName))
	regexBasketDenom = regexp.MustCompile(fmt.Sprintf(`^%s$`, RegexBasketDenom))
)

// FormatBasketDenom formats denom and display denom:
// - denom: eco.<m.Exponent><m.CreditTypeAbbrev>.<m.Name>
// - display denom: eco.<m.CreditTypeAbbrev>.<m.Name>
func FormatBasketDenom(name, creditTypeAbbrev string, exponent uint32) (string, string, error) {
	exponentPrefix, err := core.ExponentToPrefix(exponent)
	if err != nil {
		return "", "", err
	}

	denom := fmt.Sprintf("%s.%s%s.%s", denomPrefix, exponentPrefix, creditTypeAbbrev, name)
	displayDenom := fmt.Sprintf("%s.%s.%s", denomPrefix, creditTypeAbbrev, name)

	return denom, displayDenom, nil
}

// ValidateBasketName validates a basket denomination conforms to the format
// described in FormatBasketName. The return is nil if the denom is valid.
func ValidateBasketName(name string) error {
	matches := regexBasketName.FindStringSubmatch(name)
	if matches == nil {
		return fmt.Errorf("name must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long")
	}
	return nil
}

// ValidateBasketDenom validates a basket denomination conforms to the format
// described in FormatBasketDenom. The return is nil if the denom is valid.
func ValidateBasketDenom(denom string) error {
	matches := regexBasketDenom.FindStringSubmatch(denom)
	if matches == nil {
		return fmt.Errorf("%s is not a valid basket denom", denom)
	}
	return nil
}
