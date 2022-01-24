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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user data with id",
	Long:  `Get spectfic using by adding flag --id {ID} `,
	Run: func(cmd *cobra.Command, args []string) {
		getUser()

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntVarP(&user_id, "id", "i", -1, "specify id")
	getCmd.MarkFlagRequired("id")

}

func getUser() {

	res, err := http.Get(fmt.Sprintf("http://localhost:9010/users/%d", user_id))
	if err != nil {
		fmt.Println("Error While Fetching Data")
		return
	}
	defer res.Body.Close()
	if status := res.StatusCode; status != http.StatusOK {
		json.NewDecoder(res.Body).Decode(&custom_error)
		pj, _ := json.MarshalIndent(custom_error, "", "    ")
		fmt.Println(string(pj))
	} else {
		json.NewDecoder(res.Body).Decode(&user_model)
		fmt.Println("User: \n---")
		pj, _ := json.MarshalIndent(user_model, "", "    ")
		fmt.Println(string(pj))
		fmt.Println("---")
	}

}
