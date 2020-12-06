package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kkennis/nameservice/x/nameservice/keeper"
	"github.com/kkennis/nameservice/x/nameservice/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper /* TODO: Define what keepers the module needs */, data types.GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	for _, record := range data.NameRecords {
		keeper.SetName(ctx, record.Value, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	var records []types.Name
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois, _ := k.GetName(ctx, name)
		records = append(records, whois)

	}
	return types.GenesisState{NameRecords: records}
}
