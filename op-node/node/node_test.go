package node

import (
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/timeint"
	"github.com/stretchr/testify/require"
)

func TestUnixTimeStale(t *testing.T) {
	require.True(t, unixTimeStale(1_600_000_000, 1*time.Hour))
	require.False(t, unixTimeStale(timeint.FromUint64SecToSec(uint64(time.Now().Unix())), 1*time.Hour))
}
