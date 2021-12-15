package simulation

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRandomString basic determinism in randomString()
func TestRandomString(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	randString := randomStringtwo(r)
	newString := randomStringtwo(r)

	require.Equal(t, "JN", randString)
	require.Equal(t, "PGSI", newString)

}

// TestCreateTestAddresses basic determinism test
func TestCreateTestAddresses(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	addrs := createTestAddresses(r)

	require.Equal(t, []string{"enci12e0l58tn88h2ytqxl37w6g3kaz773s3m06chtg", "enci18cv6ef38zr4l25vx7rp33u7wkmkk62x60udv5m"}, addrs)

}
