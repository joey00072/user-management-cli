/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// getallCmd represents the getall command
var getallCmd = &cobra.Command{
	Use:   "getall",
	Short: "Get all users ",
	Long:  `Get Infomation of all users in details`,
	Run: func(cmd *cobra.Command, args []string) {
		getAllUsers()
	},
}

func init() {
	rootCmd.AddCommand(getallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAllUsers() {
	res, err := http.Get(fmt.Sprintf("%v", URL))

	if err != nil {
		fmt.Println("Error While Fetching Data")
		return
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&user_models)

	fmt.Println("List of all user: \n---")
	if status := res.StatusCode; status != http.StatusOK {
		json.NewDecoder(res.Body).Decode(&custom_error)
		pj, _ := json.MarshalIndent(custom_error, "", "    ")
		fmt.Println(string(pj))
	} else {
		json.NewDecoder(res.Body).Decode(&user_models)
		pj, _ := json.MarshalIndent(user_models, "", "    ")
		fmt.Println(string(pj))
	}
}
