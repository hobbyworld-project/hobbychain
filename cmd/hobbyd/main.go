package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/hobbyworld-project/hobbychain/app"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/hobbyworld-project/hobbychain/cmd/hobbyd/cmd"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cmdcfg "github.com/hobbyworld-project/hobbychain/cmd/config"
)

const (
	Version = "v0.5.2"
)

var (
	BuildTime = "2023-11-21"
	GitCommit = ""
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	version.Version = fmt.Sprintf("%s %s %s", Version, BuildTime, GitCommit)
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	// TODO fix
	// if err := cmdcfg.EnableObservability(); err != nil {
	// 	panic(err)
	// }
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
