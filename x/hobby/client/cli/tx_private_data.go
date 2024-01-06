package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPrivateData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "private-data [key] [value]",
		Short: "Broadcast message private-data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argKey := args[0]
			argValue := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPrivateData(
				clientCtx.GetFromAddress().String(),
				argKey,
				argValue,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
