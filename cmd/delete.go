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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user data with id",
	Long:  `Delete spectfic using by adding flag --id {ID} `,
	Run: func(cmd *cobra.Command, args []string) {
		deleteUser()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&user_id, "id", "i", -1, "specify id")
	deleteCmd.MarkFlagRequired("id")

}

func deleteUser() {

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%v%d", URL, user_id), nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if status := res.StatusCode; status != http.StatusOK {
		json.NewDecoder(res.Body).Decode(&custom_error)
		pj, _ := json.MarshalIndent(custom_error, "", "    ")
		fmt.Println(string(pj))
	} else {
		json.NewDecoder(res.Body).Decode(&user_model)
		fmt.Println("User Deleted: \n---")
		pj, _ := json.MarshalIndent(user_model, "", "    ")
		fmt.Println(string(pj))
		fmt.Println("---")
	}

}
