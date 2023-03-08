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
		GetStudentCmd(),
		GetAdminCmd(),
		GetAdminsCmd(),
		GetLeaveStatusCmd(),
	)

	return cmd
}

// Get leave command
func GetLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "applied-leaves",
		Short:   "Applied Leaves",
		Long:    `List of all the applied leaves by the students`,
		Example: "./simd query coslms applied-leaves",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetLeaveRequestsRequest{}
			res, err := queryClient.GetLeaveRequests(ctx.Context(), params)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Students Command
func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "students",
		Short:   "Students",
		Long:    `List all the students which are added by`,
		Example: "./simd query coslms students",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetStudentsRequest{}
			res, err := queryClient.GetStudents(ctx.Context(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Approved Leaves command
func GetApprovedLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "approved-leaves",
		Short:   "Approved Leaves",
		Long:    `List all the approved leaves which are verified by admin`,
		Example: "./simd query coslms approved-leaves",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetLeaveApprovedRequestsRequest{}
			res, err := queryClient.GetLeaveApprovedRequests(ctx.Context(), params)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Student Command
func GetStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "student  [address]",
		Short:   "Student",
		Long:    `Lists the student details based on the address input`,
		Example: "./simd query coslms student [address]",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			student := types.GetStudentRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudent(ctx.Context(), &student)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Admin Command
func GetAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "admin [address]",
		Short:   "Admin",
		Long:    `Lists the admin based on the address provided`,
		Example: "./simd query coslms admin [address]",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			admin := types.GetAdminRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetAdmin(ctx.Context(), &admin)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// Get Admins Command
func GetAdminsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "admins",
		Short:   "Admins",
		Long:    `List all the registered admins`,
		Example: "./simd query coslms admins",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetAdminsRequest{}
			res, err := queryClient.GetAdmins(ctx.Context(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func GetLeaveStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "leave-status [admin] [leaveid]",
		Short:   "Leave Status",
		Long:    `Get the leave status once the id and admin are passed`,
		Example: "./simd query coslms leave-status [admin] [leaveid]",
		RunE: func(ctx *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(ctx)
			if err != nil {
				panic(err)
			}
			status := types.GetLeaveStatusRequest{
				Admin:   args[0],
				LeaveId: args[1],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveStatus(ctx.Context(), &status)
			if err != nil {
				return err
			}
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
	rootCmd.AddCommand(GetStudentCmd())
	rootCmd.AddCommand(GetAdminsCmd())
	rootCmd.AddCommand(GetLeaveStatusCmd())
	rootCmd.AddCommand(GetAdminCmd())
	rootCmd.SuggestionsMinimumDistance = 3

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getuserCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
