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
	"time"

	"coslms/x/lms/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// txCmd represents the tx command
func GetTxCmd() *cobra.Command {
	studentTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "|lms|",
		Long:  `lms module commands`,
		RunE:  client.ValidateCmd,
	}
	studentTxCmd.AddCommand(
		RegisterAdminCmd(),
		AddStudentCmd(),
		AcceptLeaveCmd(),
		ApplyLeaveCmd(),
	)
	return studentTxCmd
	// 	Use:   "tx",
	// 	Short: "A brief description of your command",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("tx called")

}
func AddStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Adding User",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			admin := args[0]
			students := []*types.Student{}

			for i := 0; i < (len(args)-1)/3; i++ {

				student := &types.Student{
					Address: args[3*i+1],
					Name:    args[3*i+2],
					Id:      args[3*i+3],
				}
				students = append(students, student)

			}

			msgClient := types.NewAddStudentRequest(admin, students)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)

		},
	}
	return cmd
}

func RegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Admin Registration",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			address := args[0]
			name := args[1]
			msgClient := types.NewRegisterAdminRequest(address, name)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)

		},
	}
	return cmd
}

func ApplyLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Leave Application",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			var format string = "2006-Jan-06"
			fromDate, _ := time.Parse(format, args[3])
			toDate, _ := time.Parse(format, args[4])
			address := args[0]
			reason := args[1]
			leaveid := args[2]
			from := &fromDate
			to := &toDate
			msgClient := types.NewApplyLeaveRequest(address, reason, leaveid, from, to)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}
func AcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Apply Leave",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			admin := args[0]
			leaveid := args[1]
			status := args[2]
			msgClient := types.NewAcceptLeaveRequest(admin, leaveid, types.LeaveStatus_STATUS_REJECTED)
			if status == "0" {
				msgClient = types.NewAcceptLeaveRequest(admin, leaveid, types.LeaveStatus_STATUS_ACCEPTED)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}
func init() {
	rootCmd.AddCommand(AddStudentCmd())
	rootCmd.AddCommand(RegisterAdminCmd())
	rootCmd.AddCommand(ApplyLeaveCmd())
	rootCmd.AddCommand(AcceptLeaveCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
