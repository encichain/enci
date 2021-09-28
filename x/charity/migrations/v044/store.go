package v044

import (
	//"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v042charity "github.com/user/encichain/x/charity/migrations/v042"
	"github.com/user/encichain/x/charity/types"
)

func migrateTaxRate(store sdk.KVStore) {
	// Key is of format
	// prefix (0x01) || decBytes
	// Format and value does not change
	oldstore := prefix.NewStore(store, v042charity.TaxRateKey)
	oldstoreIter := oldstore.Iterator(nil, nil)
	defer oldstoreIter.Close()

	for ; oldstoreIter.Valid(); oldstoreIter.Next() {

		store.Set(types.TaxRateKey, oldstoreIter.Value())
		oldstore.Delete(oldstoreIter.Key())
	}

}

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)
	migrateTaxRate(store)
	return nil
}
