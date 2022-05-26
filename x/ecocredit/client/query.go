package client

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basketcli "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	marketplacecli "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// QueryCmd returns the parent command for all x/ecocredit query commands.
func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Args:  cobra.ExactArgs(1),
		Use:   name,
		Short: "Query commands for the ecocredit module",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		QueryClassesCmd(),
		QueryClassCmd(),
		QueryClassIssuersCmd(),
		QueryBatchesCmd(),
		QueryBatchesByIssuerCmd(),
		QueryBatchesByClassCmd(),
		QueryBatchesByProjectCmd(),
		QueryBatchCmd(),
		QueryBalanceCmd(),
		QuerySupplyCmd(),
		QueryCreditTypesCmd(),
		QueryProjectsCmd(),
		QueryProjectsByReferenceIdCmd(),
		QueryProjectCmd(),
		QueryParamsCmd(),
		basketcli.QueryBasketCmd(),
		basketcli.QueryBasketsCmd(),
		basketcli.QueryBasketBalanceCmd(),
		basketcli.QueryBasketBalancesCmd(),
		marketplacecli.QuerySellOrderCmd(),
		marketplacecli.QuerySellOrdersCmd(),
		marketplacecli.QuerySellOrdersByAddressCmd(),
		marketplacecli.QuerySellOrdersByBatchDenomCmd(),
		marketplacecli.QueryAllowedDenomsCmd(),
	)
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// QueryClassesCmd returns a query command that lists all credit classes.
func QueryClassesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes",
		Short: "List all credit classes with pagination flags",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Classes(cmd.Context(), &core.QueryClassesRequest{
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "classes")
	return qflags(cmd)
}

// QueryClassCmd returns a query command that retrieves information for a
// given credit class.
func QueryClassCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "class [class_id]",
		Short: "Retrieve credit class info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Class(cmd.Context(), &core.QueryClassRequest{
				ClassId: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryClassIssuersCmd returns a query command that retrieves addresses of the
// credit class issuers.
func QueryClassIssuersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-issuers [class-id]",
		Short: "Retrieve addresses of the issuers for a credit class",
		Long: `Retrieve addresses of the issuers for a credit class.

Args:
	class-id: credit class id
		`,
		Example: `
$ regen q ecocredit class-issuers C01
$ regen q ecocredit class-issuers C01 --pagination.limit 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ClassIssuers(cmd.Context(), &core.QueryClassIssuersRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "class-issuers")
	return qflags(cmd)
}

// QueryProjectsCmd returns a query command that retrieves projects.
func QueryProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects [class_id]",
		Short: "List all projects in the given class with pagination flags",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Projects(cmd.Context(), &core.QueryProjectsRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "projects")
	return qflags(cmd)
}

// QueryProjectCmd returns a query command that retrieves project information.
func QueryProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project [project_id]",
		Short: "Retrieve project info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.Project(cmd.Context(), &core.QueryProjectRequest{
				ProjectId: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	return qflags(cmd)
}

// QueryBatchesCmd returns a query command that retrieves credit batches for a
// given project.
func QueryBatchesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches",
		Short: "Query all credit batches with pagination flags",
		Long:  "Query all credit batches with pagination flags.",
		Example: `
regen q ecocredit batches
regen q ecocredit batches --pagination.limit 10
		`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Batches(cmd.Context(), &core.QueryBatchesRequest{
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches")

	return qflags(cmd)
}

// QueryBatchesByIssuerCmd returns a query command that retrieves credit batches based on issuer.
func QueryBatchesByIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-issuer [issuer]",
		Short: "Query all credit batches based on issuer",
		Long:  "Query all credit batches based on issuer with pagination flags.",
		Example: `
regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38
regen q ecocredit batches-by-issuer regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38 --pagination.limit 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByIssuer(cmd.Context(), &core.QueryBatchesByIssuerRequest{
				Issuer:     args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-issuer")

	return qflags(cmd)
}

// QueryBatchesByClassCmd returns a query command that retrieves credit batches for a
// given credit class.
func QueryBatchesByClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-class [class_id]",
		Short: "Query all credit batches based on credit class",
		Long:  "Query all credit batches based on credit class with pagination flags.",
		Example: `
regen q ecocredit batches-by-class C01
regen q ecocredit batches-by-class C01 --pagination.limit 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByClass(cmd.Context(), &core.QueryBatchesByClassRequest{
				ClassId:    args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-class")

	return qflags(cmd)
}

// QueryBatchesByProjectCmd returns a query command that retrieves credit batches for a
// given project.
func QueryBatchesByProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches-by-project [project_id]",
		Short: "Query all credit batches based on project",
		Long:  "Query all credit batches based on project with pagination flags.",
		Example: `
regen q ecocredit batches-by-project C01
regen q ecocredit batches-by-project C01 --pagination.limit 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.BatchesByProject(cmd.Context(), &core.QueryBatchesByProjectRequest{
				ProjectId:  args[0],
				Pagination: pagination,
			})

			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "batches-by-project")

	return qflags(cmd)
}

// QueryBatchCmd returns a query command that retrieves information for a
// given credit batch.
func QueryBatchCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "batch [batch_denom]",
		Short: "Retrieve the credit issuance batch info",
		Long:  "Retrieve the credit issuance batch info based on the batch denom.",
		Example: `
regen q ecocredit batch C01-001-20200101-20210101-001
regen q ecocredit batch C01-001-20200101-20210101-001 --pagination.limit 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.Batch(cmd.Context(), &core.QueryBatchRequest{
				BatchDenom: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryBalanceCmd returns a query command that retrieves the tradable and
// retired balances for a given credit batch and account address.
func QueryBalanceCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "balance [batch_denom] [account]",
		Short: "Retrieve the tradable and retired balances of the credit batch",
		Long:  "Retrieve the tradable and retired balances of the credit batch for a given account address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Balance(cmd.Context(), &core.QueryBalanceRequest{
				BatchDenom: args[0], Account: args[1],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QuerySupplyCmd returns a query command that retrieves the tradable and
// retired supply of credits for a given credit batch.
func QuerySupplyCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "supply [batch_denom]",
		Short: "Retrieve the tradable and retired supply of the credit batch",
		Long:  "Retrieve the tradable and retired supply of the credit batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Supply(cmd.Context(), &core.QuerySupplyRequest{
				BatchDenom: args[0],
			})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryCreditTypesCmd returns a query command that retrieves the list of
// approved credit types.
func QueryCreditTypesCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "types",
		Short: "Retrieve the list of credit types",
		Long:  "Retrieve the list of credit types that contains the type name, measurement unit and precision",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.CreditTypes(cmd.Context(), &core.QueryCreditTypesRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryParamsCmd returns ecocredit module parameters.
func QueryParamsCmd() *cobra.Command {
	return qflags(&cobra.Command{
		Use:   "params",
		Short: "Query the current ecocredit module parameters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the current ecocredit module parameters
			
Examples:
$%s query %s params
$%s q %s params
			`, version.AppName, ecocredit.ModuleName, version.AppName, ecocredit.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Params(cmd.Context(), &core.QueryParamsRequest{})
			return printQueryResponse(ctx, res, err)
		},
	})
}

// QueryProjectsByReferenceIdCmd returns command that retrieves list of projects by reference id with pagination.
func QueryProjectsByReferenceIdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects-by-reference-id [reference-id]",
		Short: "Retrieve list of projects by reference-id with pagination flags",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Retrieve list of projects by reference-id with pagination flags
			
Examples:
$%s query %s projects-by-reference-id R1
$%s q %s projects-by-reference-id R1 --limit 10
			`, version.AppName, ecocredit.ModuleName, version.AppName, ecocredit.ModuleName),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ProjectsByReferenceId(cmd.Context(), &core.QueryProjectsByReferenceIdRequest{
				ReferenceId: args[0],
				Pagination:  pagination,
			})
			return printQueryResponse(ctx, res, err)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "projects-by-reference-id")

	return qflags(cmd)
}
