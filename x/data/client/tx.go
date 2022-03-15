package client

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/data"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/data transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        name,
		Short:                      "Data transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		MsgAnchorDataCmd(),
		MsgSignDataCmd(),
		MsgDefineResolverCmd(),
		MsgRegisterResolverCmd(),
	)

	return cmd
}

// MsgAnchorDataCmd creates a CLI command for Msg/AnchorData.
func MsgAnchorDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "anchor [iri]",
		Short: "Anchors a piece of data to the blockchain based on its secure " +
			"hash, effectively providing a tamper resistant timestamp.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			iri := args[0]
			if len(iri) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("iri cannot be empty")
			}

			signer := clientCtx.GetFromAddress()
			content, err := data.ParseIRI(iri)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
			}

			msg := data.MsgAnchorData{
				Sender: signer.String(),
				Hash:   content,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgSignDataCmd creates a CLI command for Msg/SignData.
func MsgSignDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sign [iri]",
		Short:   `Sign a piece of on-chain data.`,
		Long:    `Sign a piece of on-chain data, attesting to its validity. The data MUST be of graph type (rdf file extension).`,
		Example: "regen tx data sign regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			iri := args[0]
			if len(iri) == 0 {
				return sdkerrors.ErrInvalidRequest.Wrap("iri is required")
			}

			signer := clientCtx.GetFromAddress()
			content, err := data.ParseIRI(iri)
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("invalid iri: %s", err.Error())
			}
			graph := content.GetGraph()
			if graph == nil {
				return sdkerrors.ErrInvalidRequest.Wrap("can only sign graph data types")
			}

			msg := data.MsgSignData{
				Signers: []string{signer.String()},
				Hash:    graph,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgDefineResolverCmd creates a CLI command for Msg/DefineResolver.
func MsgDefineResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "define-resolver [resolver_url]",
		Short: `Registers data content hashes.`,
		Long: `RegisterResolver registers data content hashes.
Parameters:
  resolver_url: resolver_url is a resolver URL which should refer to an HTTP service which will respond to 
			  a GET request with the IRI of a ContentHash and return the content if it exists or a 404.
Flags:
  --from: from flag is the address of the resolver manager
		`,
		Example: "regen tx data define-resolver http://foo.bar --from manager",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			resolverUrl := args[0]
			if _, err := url.ParseRequestURI(resolverUrl); err != nil {
				return err
			}

			msg := data.MsgDefineResolver{
				Manager:     clientCtx.GetFromAddress().String(),
				ResolverUrl: resolverUrl,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MsgRegisterResolverCmd creates a CLI command for Msg/RegisterResolver.
func MsgRegisterResolverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-resolver [resolver_id] [content_hashes_json]",
		Short: `registers data content hashes`,
		Long: `registers data content hashes
Parameters:
    resolver_id: resolver id is the ID of a resolver
	content_hashes_json: contains list of content hashes which the resolver claims to serve
Flags:
	--from: manager is the address of the resolver manager
		`,
		Example: `
			regen tx data register-resolver 1 content.json

			where content.json contains
			{
				"data": [
					{
						"graph": {
							"hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
							"digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
							"canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
							"merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
						}
					}
				]
			}
			`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			resolverID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid resolver id")
			}

			contentHashes, err := parseContentHashes(clientCtx, args[1])
			if err != nil {
				return err
			}

			msg := data.MsgRegisterResolver{
				Manager:    clientCtx.GetFromAddress().String(),
				ResolverId: resolverID,
				Data:       contentHashes,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseContentHashes(clientCtx client.Context, filePath string) ([]*data.ContentHash, error) {
	contentHashes := data.ContentHashes{}

	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := clientCtx.Codec.UnmarshalJSON(bz, &contentHashes); err != nil {
		return nil, err
	}

	return contentHashes.Data, nil
}
