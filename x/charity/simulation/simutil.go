package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const alphaBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//randomString generates a random string of specified size using alphaBytes
func randomString(r *rand.Rand, size int) string {
	b := strings.Builder{}
	b.Grow(size)
	for i := 0; i < size; i++ {
		b.WriteByte(alphaBytes[r.Int63()%int64(len(alphaBytes))])
	}
	return b.String()
}

// genTestAddresses generates a slice of account addresses
func genTestAddresses(amount int64) []sdk.AccAddress {
	addrs := []sdk.AccAddress{}
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")
	for i := int64(0); i < amount; i++ {
		_, _, addr := testdata.KeyTestPubAddr()
		addrs = append(addrs, addr)
	}
	return addrs
}
