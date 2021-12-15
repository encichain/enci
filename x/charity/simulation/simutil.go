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
func genTestAddresses(amount int64) []string {
	addrs := []string{}
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")
	for i := int64(0); i < amount; i++ {
		_, _, addr := testdata.KeyTestPubAddr()
		addrs = append(addrs, addr.String())
	}
	return addrs
}

//randomString generates a random string of random size using alphaBytes
func randomStringtwo(r *rand.Rand) string {
	b := strings.Builder{}
	size := r.Intn(8) + 1
	b.Grow(size)
	for i := 0; i < size; i++ {
		b.WriteByte(alphaBytes[r.Int63()%int64(len(alphaBytes))])
	}
	return b.String()
}

// createTestAddresses returns a slice of account addresses in string format from hard-coded account addresses pool to avoid determinism failure in simulation.
func createTestAddresses(r *rand.Rand) (addrs []string) {
	addrPool := []string{"enci12p43x8vny2xqpl2z6n9rx2x2gna332ee5xkc9c", "enci16ruw3nnsrt963y47y8m8h0g6p4pkyudvm5j3fc", "enci1y4xvt5y49qhqed2c2eu6xqkmu9wa356pfwsv0v",
		"enci17djxamu90muy3y2yly0ah3w5fuqug0a278h0nt", "enci1trnq89nvwen6fdhu43mw5mjl6chxf0sr8lkxqx", "enci1sjyj6s64tmdud7kshwe5nla0g0acnhmxuzpgdm",
		"enci1mm2n0kyy3g6n6hzgj8yzeqfug5t8pw7jcwvc46", "enci1qhdl3u4nze4tuhxec6yy39haz93e0gm72lk0ra", "enci1cs02xt9er8wphm4wh246q0utqagtxqkt5t0dsv",
		"enci168q953jlsfd0jgs5j27wjnqcay5mxw0q3fevut", "enci1gut4x68whaak5l9ckntr0jcasfaqjv6tlfezmf", "enci1vrvn4dvdtnncvyvlch99meyql5xs3dkng30vsk",
		"enci12e0l58tn88h2ytqxl37w6g3kaz773s3m06chtg", "enci1m9c2al4zmqhlx4vycdnq50g7mhz3c8ka20l5kq", "enci18cv6ef38zr4l25vx7rp33u7wkmkk62x60udv5m"}
	addrs = []string{}

	for i := 0; i < r.Intn(4)+1; i++ {
		addrs = append(addrs, addrPool[r.Intn(len(addrPool))])
	}

	return
}
