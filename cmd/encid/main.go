package main

import (
	"os"

	"github.com/spf13/cobra"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/encichain/enci/app"
	oraclecli "github.com/encichain/enci/x/oracle/client/cli"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"

	"github.com/encichain/enci/cmd/encid/worker"
)

func main() {
	rootCmd, _ := NewRootCmdWithWorker(
		app.AppName,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.AppName,
		app.ModuleBasics,
		app.NewEnciApp,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func NewRootCmdWithWorker(
	appName,
	accountAddressPrefix,
	defaultNodeHome,
	defaultChainID string,
	moduleBasics module.BasicManager,
	buildApp cosmoscmd.AppBuilder,
	options ...cosmoscmd.Option,
) (*cobra.Command, cosmoscmd.EncodingConfig) {
	rootCmd, enCfg := cosmoscmd.NewRootCmd(
		appName,
		accountAddressPrefix,
		defaultNodeHome,
		defaultChainID,
		moduleBasics,
		buildApp,
		options...,
	)
	oraclecli.InitializeWorker(worker.HandleBlock, worker.HandleTx)

	return rootCmd, enCfg
}
