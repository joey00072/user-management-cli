/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Name string
var Email string
var Age int
var Gender string
var Country string
var Status string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new user data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	getCmd.Flags().StringVarP(&Name, "name", "f", "", "specify name")
	getCmd.MarkFlagRequired("name")
	getCmd.Flags().StringVarP(&Email, "email", "e", "", "specify email")
	getCmd.MarkFlagRequired("email")
	getCmd.Flags().IntVarP(&Age, "age", "a", -1, "specify age")
	getCmd.MarkFlagRequired("age")
	getCmd.Flags().StringVarP(&Gender, "gender", "f", "", "specify gender")
	getCmd.MarkFlagRequired("gender")
	getCmd.Flags().StringVarP(&Status, "status", "f", "", "specify status")
	getCmd.MarkFlagRequired("status")

}
