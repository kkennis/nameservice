package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice is Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Name struct {
    Value string 			`json:"value" yaml:"value"`
	Owner sdk.AccAddress 	`json:"creator" yaml:"creator"`
    Price sdk.Coins 		`json:"price" yaml:"price"`
}

// NewName returns a new Name with the minprice as the price
func NewName Name {
	return Name{
		Price: MinNamePrice
	}
}

// implement fmt.Stringer
func (n Name) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, n.Owner, n.Value, n.Price );
}
