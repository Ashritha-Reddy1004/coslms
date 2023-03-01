/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"coslms/x/lms/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
// NewTxCmd returns a root CLI command handler for all x/lms transaction commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Query Commands",
		Long:                       "Querying commands for the LMS module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetLeavesCmd(),
		GetStudentsCmd(),
		GetApprovedLeavesCmd(),
	)

	return cmd
}

// Get leave command
func GetLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getleaves",
		Short: "List all applied leaves",
		Long:  `List all the leaves which are applied by the students`,
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetLeaveRequestsRequest{}
			res, err := queryClient.GetLeaveRequests(ctx.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Students Command
func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstudents",
		Short: "List all the students",
		Long:  `List all the students which are added by admin`,
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetStudentsRequest{}
			res, err := queryClient.GetStudents(ctx.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Approved Leaves command
func GetApprovedLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getapprovedleaves",
		Short: "List all approved leaves by admin",
		Long:  `List all the approved leaves which are verified by admin`,
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetLeaveApprovedRequestsRequest{}
			res, err := queryClient.GetLeaveApprovedRequests(ctx.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func init() {
	rootCmd.AddCommand(GetStudentsCmd())
	rootCmd.AddCommand(GetLeavesCmd())
	rootCmd.AddCommand(GetApprovedLeavesCmd())
	// rootCmd.SuggestionsMinimumDistance = 3

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getuserCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
