package worker

import (
	"github.com/spf13/cobra"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleTx is our custom tx handler
func HandleTx(cmd *cobra.Command, txEvent rpctypes.ResultEvent) error {
	return nil
}

// HandleBlock is our custom block handler
func HandleBlock(cmd *cobra.Command, blockEvent rpctypes.ResultEvent) error {
	return nil
}
