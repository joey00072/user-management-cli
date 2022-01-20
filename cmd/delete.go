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
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	getCmd.Flags().IntVarP(&id, "id", "f", -1, "specify id")
	getCmd.MarkFlagRequired("id")

}

func deleteUser(id int) {

	fmt.Println(id)

	res, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:9010/users/%d", id), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&user)

	fmt.Println("User: \n---")

	fmt.Printf("id : %v\n", user.Id)
	fmt.Printf("name : %v\n", user.Name)
	fmt.Printf("email : %v\n", user.Email)
	fmt.Printf("age : %v\n", user.Age)
	fmt.Printf("gender : %v\n", user.Gender)
	fmt.Printf("country : %v\n", user.Country)
	fmt.Printf("status : %v\n", user.Status)
	fmt.Println("---")
}
